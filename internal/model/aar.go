package model

import (
	"time"

	"github.com/google/uuid"
)

// HuntAfterActionReport represents post-hunt documentation with narrative and images.
// Each hunt has at most one AAR (1:1 relationship enforced by unique constraint on HuntID).
type HuntAfterActionReport struct {
	ID          uuid.UUID
	HuntID      uuid.UUID
	Description string
	ImageURLs   []string
	CreatedByID uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// AARParticipant represents a user tagged in an after-action report.
type AARParticipant struct {
	AARID  uuid.UUID
	UserID uuid.UUID
}
