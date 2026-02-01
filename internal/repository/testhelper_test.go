package repository_test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// testDB opens a connection to the test database.
// Tests using this must have DATABASE_URL set and a running PostgreSQL instance.
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

// withTestTx starts a transaction, runs fn, then rolls back.
// This provides test isolation without mutating the database.
func withTestTx(t *testing.T, db *sql.DB, fn func(tx *sql.Tx)) {
	t.Helper()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("beginning transaction: %v", err)
	}
	defer func() { _ = tx.Rollback() }()

	fn(tx)
}
