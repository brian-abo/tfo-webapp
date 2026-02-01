package model

import (
	"time"

	"github.com/google/uuid"
)

// ContactSubmission represents a message submitted through the contact form.
type ContactSubmission struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Message   string
	CreatedAt time.Time
}
