package model

import (
	"time"

	"github.com/google/uuid"
)

// LotteryResult represents the outcome of a lottery draw for a signup.
// Each signup has at most one result (unique constraint on SignupID).
// Results are immutable; acceptance/decline is tracked in the Confirmation model.
type LotteryResult struct {
	ID               uuid.UUID
	HuntID           uuid.UUID
	SignupID         uuid.UUID
	Position         int
	AuditSeed        int64
	AlgorithmVersion string
	DrawnAt          time.Time
	CreatedAt        time.Time
}

// IsPrimary returns true if this result is within the primary capacity.
// Requires the hunt's primary capacity to determine the cutoff.
func (r *LotteryResult) IsPrimary(primaryCapacity int) bool {
	return r.Position <= primaryCapacity
}
