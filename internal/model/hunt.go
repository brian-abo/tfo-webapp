package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// HuntStatus represents the lifecycle state of a hunt.
type HuntStatus string

const (
	HuntStatusDraft     HuntStatus = "draft"
	HuntStatusOpen      HuntStatus = "open"
	HuntStatusClosed    HuntStatus = "closed"
	HuntStatusCompleted HuntStatus = "completed"
	HuntStatusCancelled HuntStatus = "cancelled"
)

// Hunt represents a hunting event that members can sign up for.
type Hunt struct {
	ID                uuid.UUID
	Title             string
	Description       string
	Location          string
	ImageURLs         []string
	Qualifiers        sql.NullString
	HuntDate          time.Time
	SignupWindowStart time.Time
	SignupWindowEnd   time.Time
	PrimaryCapacity   int
	AlternateCapacity int
	Status            HuntStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// TotalCapacity returns the combined primary and alternate capacity.
func (h *Hunt) TotalCapacity() int {
	return h.PrimaryCapacity + h.AlternateCapacity
}

// IsSignupWindowOpen returns true if the current time is within the signup window.
func (h *Hunt) IsSignupWindowOpen(now time.Time) bool {
	return now.After(h.SignupWindowStart) && now.Before(h.SignupWindowEnd)
}

// CanAcceptSignups returns true if the hunt is open and within the signup window.
func (h *Hunt) CanAcceptSignups(now time.Time) bool {
	return h.Status == HuntStatusOpen && h.IsSignupWindowOpen(now)
}
