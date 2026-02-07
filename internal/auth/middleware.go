package auth

import (
	"log"
	"net/http"

	"github.com/brian-abo/tfo-webapp/internal/model"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

const (
	// SessionCookieName is the name of the session cookie.
	SessionCookieName = "tfo_session"
)

// Middleware provides authentication middleware functions.
type Middleware struct {
	sessions SessionStore
	users    *repository.UserRepository
}

// NewMiddleware creates a new auth middleware.
func NewMiddleware(sessions SessionStore, users *repository.UserRepository) *Middleware {
	return &Middleware{
		sessions: sessions,
		users:    users,
	}
}

// LoadSession is middleware that loads the session and user from the cookie.
// It attaches them to the request context for downstream handlers.
func (m *Middleware) LoadSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(SessionCookieName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		session, err := m.sessions.Get(r.Context(), cookie.Value)
		if err != nil {
			log.Printf("error loading session: %v", err)
			next.ServeHTTP(w, r)
			return
		}
		if session == nil {
			// Session expired or invalid, clear the cookie
			clearSessionCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		// Touch session to extend expiry (sliding window)
		if err := m.sessions.Touch(r.Context(), session.ID); err != nil {
			log.Printf("error touching session: %v", err)
		}

		// Load user
		user, err := m.users.FindByID(r.Context(), session.UserID)
		if err != nil {
			log.Printf("error loading user: %v", err)
			next.ServeHTTP(w, r)
			return
		}
		if user == nil {
			// User deleted, clear session
			_ = m.sessions.Delete(r.Context(), session.ID)
			clearSessionCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		// Attach session and user to context
		ctx := WithSession(r.Context(), session)
		ctx = WithUser(ctx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireAuth is middleware that requires a valid session.
// Redirects to login page if not authenticated.
func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r.Context()) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireRole returns middleware that requires the user to have a specific role.
// Returns 403 Forbidden if the user lacks the required role.
func (m *Middleware) RequireRole(role model.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := GetUser(r.Context())
			if user == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			if !hasRole(user.Role, role) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RequireAdmin is middleware that requires admin role.
func (m *Middleware) RequireAdmin(next http.Handler) http.Handler {
	return m.RequireRole(model.RoleAdmin)(next)
}

// RequireStaff is middleware that requires staff or admin role.
func (m *Middleware) RequireStaff(next http.Handler) http.Handler {
	return m.RequireRole(model.RoleStaff)(next)
}

// hasRole checks if userRole has at least the required role level.
// Hierarchy: admin > staff > member
func hasRole(userRole model.Role, required model.Role) bool {
	roleLevel := map[model.Role]int{
		model.RoleMember: 1,
		model.RoleStaff:  2,
		model.RoleAdmin:  3,
	}
	return roleLevel[userRole] >= roleLevel[required]
}

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
