// Package model defines domain entities for the TFO webapp.
package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Role represents a user's permission level in the system.
type Role string

const (
	RoleMember Role = "member"
	RoleStaff  Role = "staff"
	RoleAdmin  Role = "admin"
)

// MembershipStatus represents a user's account status.
type MembershipStatus string

const (
	MembershipPending   MembershipStatus = "pending"
	MembershipActive    MembershipStatus = "active"
	MembershipInactive  MembershipStatus = "inactive"
	MembershipSuspended MembershipStatus = "suspended"
)

// User represents an authenticated user in the system.
type User struct {
	ID               uuid.UUID
	Email            string
	Name             string
	Phone            sql.NullString
	BranchOfService  string
	Role             Role
	MembershipStatus MembershipStatus
	FacebookID       sql.NullString
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
}

// IsDeleted returns true if the user has been soft-deleted.
func (u *User) IsDeleted() bool {
	return u.DeletedAt.Valid
}

// IsActive returns true if the user's membership is active and not deleted.
func (u *User) IsActive() bool {
	return u.MembershipStatus == MembershipActive && !u.IsDeleted()
}

// IsAccountComplete returns true if required profile fields are populated.
func (u *User) IsAccountComplete() bool {
	return u.BranchOfService != ""
}
