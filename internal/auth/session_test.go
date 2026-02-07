package auth_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/onsi/gomega"

	"github.com/brian-abo/tfo-webapp/internal/auth"
	"github.com/brian-abo/tfo-webapp/internal/model"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

func testDB(t *testing.T) *sql.DB {
	t.Helper()

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		t.Skip("DATABASE_URL not set, skipping integration test")
	}

	db, err := sql.Open("pgx", url)
	if err != nil {
		t.Fatalf("opening test database: %v", err)
	}

	if err := db.PingContext(t.Context()); err != nil {
		t.Fatalf("pinging test database: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("closing test database: %v", err)
		}
	})
	return db
}

func withTestTx(t *testing.T, db *sql.DB, fn func(tx *sql.Tx)) {
	t.Helper()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("beginning transaction: %v", err)
	}
	defer func() { _ = tx.Rollback() }()

	fn(tx)
}

func createTestUser(t *testing.T, tx *sql.Tx) *model.User {
	t.Helper()
	g := NewWithT(t)

	repo := repository.NewUserRepository(tx)
	user := &model.User{
		Email:            "session-test-" + uuid.New().String() + "@example.com",
		Name:             "Session Test User",
		BranchOfService:  "Army",
		Role:             model.RoleMember,
		MembershipStatus: model.MembershipActive,
	}
	created, err := repo.Create(t.Context(), user)
	g.Expect(err).ToNot(HaveOccurred())
	return created
}

func TestPostgresSessionStore_CreateAndGet(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		user := createTestUser(t, tx)

		// Create a wrapper that implements the session store interface with tx
		store := &txSessionStore{tx: tx}

		session, token, err := store.Create(t.Context(), user.ID)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(session).ToNot(BeNil())
		g.Expect(token).ToNot(BeEmpty())
		g.Expect(session.UserID).To(Equal(user.ID))
		g.Expect(session.CSRFToken).ToNot(BeEmpty())
		g.Expect(session.ExpiresAt.After(time.Now())).To(BeTrue())

		// Get session by token
		found, err := store.Get(t.Context(), token)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).ToNot(BeNil())
		g.Expect(found.ID).To(Equal(session.ID))
	})
}

func TestPostgresSessionStore_Delete(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		user := createTestUser(t, tx)
		store := &txSessionStore{tx: tx}

		session, token, err := store.Create(t.Context(), user.ID)
		g.Expect(err).ToNot(HaveOccurred())

		err = store.Delete(t.Context(), session.ID)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := store.Get(t.Context(), token)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).To(BeNil())
	})
}

func TestPostgresSessionStore_Touch(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		user := createTestUser(t, tx)
		store := &txSessionStore{tx: tx}

		session, token, err := store.Create(t.Context(), user.ID)
		g.Expect(err).ToNot(HaveOccurred())

		originalLastActive := session.LastActiveAt

		// Touch the session
		time.Sleep(10 * time.Millisecond)
		err = store.Touch(t.Context(), session.ID)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := store.Get(t.Context(), token)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found.LastActiveAt.After(originalLastActive)).To(BeTrue())
	})
}

// txSessionStore wraps a transaction to implement session store operations for testing.
type txSessionStore struct {
	tx *sql.Tx
}

func (s *txSessionStore) Create(ctx context.Context, userID uuid.UUID) (*auth.Session, string, error) {
	token, tokenHash, err := auth.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	csrfToken, err := auth.GenerateCSRFToken()
	if err != nil {
		return nil, "", err
	}

	now := time.Now()
	expiresAt := now.Add(auth.SessionDuration)

	var session auth.Session
	err = s.tx.QueryRowContext(ctx,
		`INSERT INTO sessions (user_id, token_hash, csrf_token, created_at, expires_at, last_active_at)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, user_id, token_hash, csrf_token, created_at, expires_at, last_active_at`,
		userID, tokenHash, csrfToken, now, expiresAt, now,
	).Scan(&session.ID, &session.UserID, &session.TokenHash, &session.CSRFToken,
		&session.CreatedAt, &session.ExpiresAt, &session.LastActiveAt)
	if err != nil {
		return nil, "", err
	}

	return &session, token, nil
}

func (s *txSessionStore) Get(ctx context.Context, token string) (*auth.Session, error) {
	tokenHash := auth.HashToken(token)

	var session auth.Session
	err := s.tx.QueryRowContext(ctx,
		`SELECT id, user_id, token_hash, csrf_token, created_at, expires_at, last_active_at
		 FROM sessions
		 WHERE token_hash = $1`,
		tokenHash,
	).Scan(&session.ID, &session.UserID, &session.TokenHash, &session.CSRFToken,
		&session.CreatedAt, &session.ExpiresAt, &session.LastActiveAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if session.IsExpired() {
		return nil, nil
	}

	return &session, nil
}

func (s *txSessionStore) Touch(ctx context.Context, sessionID uuid.UUID) error {
	now := time.Now()
	newExpiry := now.Add(auth.SessionDuration)

	_, err := s.tx.ExecContext(ctx,
		`UPDATE sessions SET last_active_at = $1, expires_at = $2 WHERE id = $3`,
		now, newExpiry, sessionID,
	)
	return err
}

func (s *txSessionStore) Delete(ctx context.Context, sessionID uuid.UUID) error {
	_, err := s.tx.ExecContext(ctx,
		`DELETE FROM sessions WHERE id = $1`,
		sessionID,
	)
	return err
}

func (s *txSessionStore) DeleteExpired(ctx context.Context) (int64, error) {
	return 0, nil
}

func (s *txSessionStore) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	return nil
}

var _ auth.SessionStore = (*txSessionStore)(nil)
