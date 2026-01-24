// Package repository provides database access patterns and interfaces.
package repository

import (
	"context"
	"database/sql"
)

// DBTX is an interface that abstracts database operations, allowing
// repositories to work with both *sql.DB and *sql.Tx. This enables
// transaction support without changing repository method signatures.
type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Compile-time assertions to verify that *sql.DB and *sql.Tx implement DBTX.
var (
	_ DBTX = (*sql.DB)(nil)
	_ DBTX = (*sql.Tx)(nil)
)
