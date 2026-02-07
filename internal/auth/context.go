package auth

import (
	"context"

	"github.com/brian-abo/tfo-webapp/internal/model"
)

type contextKey string

const (
	sessionKey contextKey = "session"
	userKey    contextKey = "user"
)

// WithSession returns a new context with the session attached.
func WithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

// GetSession retrieves the session from the context.
// Returns nil if no session is present.
func GetSession(ctx context.Context) *Session {
	session, _ := ctx.Value(sessionKey).(*Session)
	return session
}

// WithUser returns a new context with the user attached.
func WithUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser retrieves the user from the context.
// Returns nil if no user is present.
func GetUser(ctx context.Context) *model.User {
	user, _ := ctx.Value(userKey).(*model.User)
	return user
}

// IsAuthenticated returns true if there is a valid user in the context.
func IsAuthenticated(ctx context.Context) bool {
	return GetUser(ctx) != nil
}

// GetCSRFToken returns the CSRF token for the current session.
// Returns empty string if no session is present.
func GetCSRFToken(ctx context.Context) string {
	session := GetSession(ctx)
	if session == nil {
		return ""
	}
	return session.CSRFToken
}
