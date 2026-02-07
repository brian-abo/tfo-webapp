// Package auth provides authentication and session management.
package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// SessionDuration is the maximum lifetime of a session.
	SessionDuration = 20 * time.Minute

	// SessionIdleTimeout is how long a session can be idle before expiring.
	SessionIdleTimeout = 15 * time.Minute

	// TokenLength is the number of random bytes in a session token.
	TokenLength = 32
)

// Session represents an authenticated user session.
type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	TokenHash    string
	CSRFToken    string
	CreatedAt    time.Time
	ExpiresAt    time.Time
	LastActiveAt time.Time
}

// IsExpired returns true if the session has expired.
func (s *Session) IsExpired() bool {
	now := time.Now()
	return now.After(s.ExpiresAt) || now.After(s.LastActiveAt.Add(SessionIdleTimeout))
}

// SessionStore defines the interface for session persistence.
// Implementations can use PostgreSQL, Redis, or other backends.
type SessionStore interface {
	// Create stores a new session and returns it.
	Create(ctx context.Context, userID uuid.UUID) (*Session, string, error)

	// Get retrieves a session by its token. Returns nil if not found or expired.
	Get(ctx context.Context, token string) (*Session, error)

	// Touch updates the last_active_at timestamp to extend the session.
	Touch(ctx context.Context, sessionID uuid.UUID) error

	// Delete removes a session.
	Delete(ctx context.Context, sessionID uuid.UUID) error

	// DeleteExpired removes all expired sessions.
	DeleteExpired(ctx context.Context) (int64, error)

	// DeleteByUserID removes all sessions for a user.
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
}

// GenerateToken creates a cryptographically secure random token.
// Returns the raw token (for cookie) and its hash (for storage).
func GenerateToken() (token string, hash string, err error) {
	bytes := make([]byte, TokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", fmt.Errorf("generating random bytes: %w", err)
	}

	token = base64.URLEncoding.EncodeToString(bytes)
	hash = HashToken(token)
	return token, hash, nil
}

// HashToken computes a SHA-256 hash of a token for secure storage.
func HashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(h[:])
}

// GenerateCSRFToken creates a random CSRF token.
func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, TokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generating CSRF token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
