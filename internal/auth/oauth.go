package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/brian-abo/tfo-webapp/internal/config"
	"github.com/brian-abo/tfo-webapp/internal/model"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

const (
	facebookAuthURL  = "https://www.facebook.com/v18.0/dialog/oauth"
	facebookTokenURL = "https://graph.facebook.com/v18.0/oauth/access_token"
	facebookGraphURL = "https://graph.facebook.com/v18.0/me"

	oauthStateCookieName = "tfo_oauth_state"
)

// OAuthHandler handles Facebook OAuth authentication.
type OAuthHandler struct {
	cfg      config.FacebookConfig
	baseURL  string
	sessions SessionStore
	users    *repository.UserRepository
}

// NewOAuthHandler creates a new OAuth handler.
func NewOAuthHandler(cfg config.FacebookConfig, baseURL string, sessions SessionStore, users *repository.UserRepository) *OAuthHandler {
	return &OAuthHandler{
		cfg:      cfg,
		baseURL:  baseURL,
		sessions: sessions,
		users:    users,
	}
}

// HandleLogin redirects the user to Facebook for authentication.
func (h *OAuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	state, err := generateState()
	if err != nil {
		log.Printf("error generating OAuth state: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Store state in cookie for CSRF protection
	http.SetCookie(w, &http.Cookie{
		Name:     oauthStateCookieName,
		Value:    state,
		Path:     "/",
		MaxAge:   600, // 10 minutes
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	redirectURI := h.baseURL + "/auth/facebook/callback"
	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&state=%s&scope=email,public_profile",
		facebookAuthURL,
		url.QueryEscape(h.cfg.ClientID),
		url.QueryEscape(redirectURI),
		url.QueryEscape(state),
	)

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// HandleCallback handles the OAuth callback from Facebook.
func (h *OAuthHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Verify state to prevent CSRF
	stateCookie, err := r.Cookie(oauthStateCookieName)
	if err != nil {
		http.Error(w, "Missing state cookie", http.StatusBadRequest)
		return
	}

	state := r.URL.Query().Get("state")
	if state != stateCookie.Value {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:   oauthStateCookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Check for error from Facebook
	if errCode := r.URL.Query().Get("error"); errCode != "" {
		errDesc := r.URL.Query().Get("error_description")
		log.Printf("OAuth error: %s - %s", errCode, errDesc)
		http.Redirect(w, r, "/login?error=oauth_failed", http.StatusSeeOther)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	// Exchange code for access token
	token, err := h.exchangeCode(r.Context(), code)
	if err != nil {
		log.Printf("error exchanging code: %v", err)
		http.Redirect(w, r, "/login?error=oauth_failed", http.StatusSeeOther)
		return
	}

	// Get user info from Facebook
	fbUser, err := h.getFacebookUser(r.Context(), token)
	if err != nil {
		log.Printf("error getting Facebook user: %v", err)
		http.Redirect(w, r, "/login?error=oauth_failed", http.StatusSeeOther)
		return
	}

	// Find or create user
	user, err := h.findOrCreateUser(r.Context(), fbUser)
	if err != nil {
		log.Printf("error finding/creating user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create session
	session, sessionToken, err := h.sessions.Create(r.Context(), user.ID)
	if err != nil {
		log.Printf("error creating session: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionToken,
		Path:     "/",
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect to dashboard or home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HandleLogout destroys the session and clears the cookie.
func (h *OAuthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session := GetSession(r.Context())
	if session != nil {
		if err := h.sessions.Delete(r.Context(), session.ID); err != nil {
			log.Printf("error deleting session: %v", err)
		}
	}

	clearSessionCookie(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type facebookUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *OAuthHandler) exchangeCode(ctx context.Context, code string) (string, error) {
	redirectURI := h.baseURL + "/auth/facebook/callback"

	data := url.Values{}
	data.Set("client_id", h.cfg.ClientID)
	data.Set("client_secret", h.cfg.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, facebookTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("creating token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("exchanging code: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("token exchange failed: %s - %s", resp.Status, string(body))
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decoding token response: %w", err)
	}

	return result.AccessToken, nil
}

func (h *OAuthHandler) getFacebookUser(ctx context.Context, accessToken string) (*facebookUser, error) {
	reqURL := fmt.Sprintf("%s?fields=id,name,email&access_token=%s", facebookGraphURL, url.QueryEscape(accessToken))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating user request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching user: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user fetch failed: %s - %s", resp.Status, string(body))
	}

	var user facebookUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding user response: %w", err)
	}

	return &user, nil
}

func (h *OAuthHandler) findOrCreateUser(ctx context.Context, fbUser *facebookUser) (*model.User, error) {
	// Try to find by Facebook ID first
	user, err := h.users.FindByFacebookID(ctx, fbUser.ID)
	if err != nil {
		return nil, fmt.Errorf("finding user by Facebook ID: %w", err)
	}
	if user != nil {
		return user, nil
	}

	// Try to find by email and link Facebook ID
	if fbUser.Email != "" {
		user, err = h.users.FindByEmail(ctx, fbUser.Email)
		if err != nil {
			return nil, fmt.Errorf("finding user by email: %w", err)
		}
		if user != nil {
			user.FacebookID = sql.NullString{String: fbUser.ID, Valid: true}
			if err := h.users.Update(ctx, user); err != nil {
				return nil, fmt.Errorf("linking Facebook ID: %w", err)
			}
			return user, nil
		}
	}

	// Create new user
	newUser := &model.User{
		Email:            fbUser.Email,
		Name:             fbUser.Name,
		BranchOfService:  "", // Will be collected later
		Role:             model.RoleMember,
		MembershipStatus: model.MembershipPending,
		FacebookID:       sql.NullString{String: fbUser.ID, Valid: true},
	}

	created, err := h.users.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	return created, nil
}

func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
