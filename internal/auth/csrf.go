package auth

import (
	"crypto/subtle"
	"net/http"
)

const (
	// CSRFTokenHeader is the HTTP header name for CSRF tokens.
	CSRFTokenHeader = "X-CSRF-Token"

	// CSRFTokenFormField is the form field name for CSRF tokens.
	CSRFTokenFormField = "csrf_token"
)

// RequireCSRF is middleware that validates CSRF tokens on state-changing requests.
// It checks for the token in the X-CSRF-Token header or csrf_token form field.
func RequireCSRF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only validate on state-changing methods
		if r.Method == http.MethodGet || r.Method == http.MethodHead || r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		session := GetSession(r.Context())
		if session == nil {
			http.Error(w, "Forbidden: no session", http.StatusForbidden)
			return
		}

		// Check header first, then form field
		token := r.Header.Get(CSRFTokenHeader)
		if token == "" {
			token = r.FormValue(CSRFTokenFormField)
		}

		if token == "" {
			http.Error(w, "Forbidden: missing CSRF token", http.StatusForbidden)
			return
		}

		// Constant-time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(token), []byte(session.CSRFToken)) != 1 {
			http.Error(w, "Forbidden: invalid CSRF token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
