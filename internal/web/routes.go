package web

import (
	"database/sql"
	"net/http"

	"github.com/brian-abo/tfo-webapp/internal/auth"
	"github.com/brian-abo/tfo-webapp/internal/handler/about"
	contactHandler "github.com/brian-abo/tfo-webapp/internal/handler/contact"
	"github.com/brian-abo/tfo-webapp/internal/handler/gallery"
	"github.com/brian-abo/tfo-webapp/internal/handler/home"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

// NewRouter creates and configures the HTTP router.
func NewRouter(db *sql.DB, authMiddleware *auth.Middleware, oauth *auth.OAuthHandler) http.Handler {
	mux := http.NewServeMux()

	// Repositories
	contactRepo := repository.NewContactRepository(db)

	// Handlers
	contact := contactHandler.NewHandler(contactRepo)

	// Static assets (no auth required)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Public pages
	mux.HandleFunc("GET /", home.Index)
	mux.HandleFunc("GET /about", about.Index)
	mux.HandleFunc("GET /gallery", gallery.Index)
	mux.HandleFunc("GET /contact", contact.Index)
	mux.HandleFunc("POST /contact", contact.Submit)

	// Auth routes
	mux.HandleFunc("GET /login", loginPage(oauth != nil))
	if oauth != nil {
		mux.HandleFunc("GET /auth/facebook", oauth.HandleLogin)
		mux.HandleFunc("GET /auth/facebook/callback", oauth.HandleCallback)
		mux.HandleFunc("POST /auth/logout", oauth.HandleLogout)
	}

	// Apply session loading middleware to all routes
	return authMiddleware.LoadSession(mux)
}

// loginPage renders the login page.
// TODO: Move to proper handler/template when login UI is implemented.
func loginPage(oauthEnabled bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if oauthEnabled {
			_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head><title>Login - The Fallen Outdoors</title></head>
<body>
<h1>Login</h1>
<p><a href="/auth/facebook">Login with Facebook</a></p>
</body>
</html>`))
		} else {
			_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head><title>Login - The Fallen Outdoors</title></head>
<body>
<h1>Login</h1>
<p>Login is not configured. Please set FACEBOOK_CLIENT_ID and FACEBOOK_CLIENT_SECRET.</p>
</body>
</html>`))
		}
	}
}
