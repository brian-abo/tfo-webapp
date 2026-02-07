package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// PostgresSessionStore implements SessionStore using PostgreSQL.
type PostgresSessionStore struct {
	db *sql.DB
}

// NewPostgresSessionStore creates a new PostgreSQL-backed session store.
func NewPostgresSessionStore(db *sql.DB) *PostgresSessionStore {
	return &PostgresSessionStore{db: db}
}

// Create stores a new session and returns it along with the raw token.
func (s *PostgresSessionStore) Create(ctx context.Context, userID uuid.UUID) (*Session, string, error) {
	token, tokenHash, err := GenerateToken()
	if err != nil {
		return nil, "", fmt.Errorf("generating token: %w", err)
	}

	csrfToken, err := GenerateCSRFToken()
	if err != nil {
		return nil, "", fmt.Errorf("generating CSRF token: %w", err)
	}

	now := time.Now()
	expiresAt := now.Add(SessionDuration)

	var session Session
	err = s.db.QueryRowContext(ctx,
		`INSERT INTO sessions (user_id, token_hash, csrf_token, created_at, expires_at, last_active_at)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, user_id, token_hash, csrf_token, created_at, expires_at, last_active_at`,
		userID, tokenHash, csrfToken, now, expiresAt, now,
	).Scan(&session.ID, &session.UserID, &session.TokenHash, &session.CSRFToken,
		&session.CreatedAt, &session.ExpiresAt, &session.LastActiveAt)
	if err != nil {
		return nil, "", fmt.Errorf("inserting session: %w", err)
	}

	return &session, token, nil
}

// Get retrieves a session by its token. Returns nil if not found or expired.
func (s *PostgresSessionStore) Get(ctx context.Context, token string) (*Session, error) {
	tokenHash := HashToken(token)

	var session Session
	err := s.db.QueryRowContext(ctx,
		`SELECT id, user_id, token_hash, csrf_token, created_at, expires_at, last_active_at
		 FROM sessions
		 WHERE token_hash = $1`,
		tokenHash,
	).Scan(&session.ID, &session.UserID, &session.TokenHash, &session.CSRFToken,
		&session.CreatedAt, &session.ExpiresAt, &session.LastActiveAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying session: %w", err)
	}

	if session.IsExpired() {
		// Clean up expired session
		_ = s.Delete(ctx, session.ID)
		return nil, nil
	}

	return &session, nil
}

// Touch updates the last_active_at timestamp to extend the session.
func (s *PostgresSessionStore) Touch(ctx context.Context, sessionID uuid.UUID) error {
	now := time.Now()
	newExpiry := now.Add(SessionDuration)

	_, err := s.db.ExecContext(ctx,
		`UPDATE sessions SET last_active_at = $1, expires_at = $2 WHERE id = $3`,
		now, newExpiry, sessionID,
	)
	if err != nil {
		return fmt.Errorf("touching session: %w", err)
	}
	return nil
}

// Delete removes a session.
func (s *PostgresSessionStore) Delete(ctx context.Context, sessionID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx,
		`DELETE FROM sessions WHERE id = $1`,
		sessionID,
	)
	if err != nil {
		return fmt.Errorf("deleting session: %w", err)
	}
	return nil
}

// DeleteExpired removes all expired sessions.
func (s *PostgresSessionStore) DeleteExpired(ctx context.Context) (int64, error) {
	now := time.Now()
	idleThreshold := now.Add(-SessionIdleTimeout)

	result, err := s.db.ExecContext(ctx,
		`DELETE FROM sessions WHERE expires_at < $1 OR last_active_at < $2`,
		now, idleThreshold,
	)
	if err != nil {
		return 0, fmt.Errorf("deleting expired sessions: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("getting rows affected: %w", err)
	}
	return count, nil
}

// DeleteByUserID removes all sessions for a user.
func (s *PostgresSessionStore) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx,
		`DELETE FROM sessions WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return fmt.Errorf("deleting user sessions: %w", err)
	}
	return nil
}

// Compile-time check that PostgresSessionStore implements SessionStore.
var _ SessionStore = (*PostgresSessionStore)(nil)
