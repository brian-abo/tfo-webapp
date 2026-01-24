package repository

import (
	"context"
	"database/sql"
	"testing"
)

// WithTestTx runs a test function within a database transaction that is
// rolled back after the test completes. This ensures test isolation without
// leaving test data in the database.
//
// Usage:
//
//	func TestSomething(t *testing.T) {
//	    repository.WithTestTx(t, db, func(tx *sql.Tx) {
//	        // Use tx for database operations
//	        // Transaction automatically rolls back when function returns
//	    })
//	}
func WithTestTx(t *testing.T, db *sql.DB, fn func(tx *sql.Tx)) {
	t.Helper()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			t.Errorf("failed to rollback transaction: %v", err)
		}
	}()

	fn(tx)
}
