package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Signup represents a user's lottery entry for a hunt.
// Signups are immutable once created; withdrawal is tracked via WithdrawnAt.
type Signup struct {
	ID                  uuid.UUID
	UserID              uuid.UUID
	HuntID              uuid.UUID
	EligibilitySnapshot sql.NullString
	CreatedAt           time.Time
	WithdrawnAt         sql.NullTime
}

// IsWithdrawn returns true if the user has withdrawn from this signup.
func (s *Signup) IsWithdrawn() bool {
	return s.WithdrawnAt.Valid
}
