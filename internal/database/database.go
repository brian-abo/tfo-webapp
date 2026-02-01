// Package database provides PostgreSQL connection management.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
)

const (
	maxRetries    = 5
	baseDelay     = 1 * time.Second
	maxDelay      = 15 * time.Second
	pingTimeout   = 5 * time.Second
	backoffFactor = 2
)

// Connect opens a PostgreSQL connection pool and verifies connectivity
// with exponential backoff. The provided context controls the overall
// connection attempt and can be used for cancellation. Returns an error
// if all retries are exhausted or the context is cancelled.
func Connect(ctx context.Context, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	delay := baseDelay
	for attempt := 1; attempt <= maxRetries; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, pingTimeout)
		err = db.PingContext(pingCtx)
		cancel()

		if err == nil {
			return db, nil
		}

		if ctx.Err() != nil {
			break
		}

		log.Printf("database connection attempt %d/%d failed: %v", attempt, maxRetries, err)

		if attempt < maxRetries {
			log.Printf("retrying in %s...", delay)
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				break
			}
			delay *= backoffFactor
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}

	if closeErr := db.Close(); closeErr != nil {
		log.Printf("closing database after failed connection: %v", closeErr)
	}
	return nil, fmt.Errorf("database connection failed after %d attempts: %w", maxRetries, err)
}
