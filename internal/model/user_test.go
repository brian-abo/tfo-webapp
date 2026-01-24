package model

import (
	"database/sql"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestUser_IsDeleted(t *testing.T) {
	t.Run("returns false when DeletedAt is null", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{DeletedAt: sql.NullTime{Valid: false}}

		g.Expect(u.IsDeleted()).To(BeFalse())
	})

	t.Run("returns true when DeletedAt is set", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{DeletedAt: sql.NullTime{Time: time.Now(), Valid: true}}

		g.Expect(u.IsDeleted()).To(BeTrue())
	})
}

func TestUser_IsActive(t *testing.T) {
	t.Run("returns true when active and not deleted", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{
			MembershipStatus: MembershipActive,
			DeletedAt:        sql.NullTime{Valid: false},
		}

		g.Expect(u.IsActive()).To(BeTrue())
	})

	t.Run("returns false when active but deleted", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{
			MembershipStatus: MembershipActive,
			DeletedAt:        sql.NullTime{Time: time.Now(), Valid: true},
		}

		g.Expect(u.IsActive()).To(BeFalse())
	})

	t.Run("returns false when pending", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{
			MembershipStatus: MembershipPending,
			DeletedAt:        sql.NullTime{Valid: false},
		}

		g.Expect(u.IsActive()).To(BeFalse())
	})

	t.Run("returns false when inactive", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{
			MembershipStatus: MembershipInactive,
			DeletedAt:        sql.NullTime{Valid: false},
		}

		g.Expect(u.IsActive()).To(BeFalse())
	})

	t.Run("returns false when suspended", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{
			MembershipStatus: MembershipSuspended,
			DeletedAt:        sql.NullTime{Valid: false},
		}

		g.Expect(u.IsActive()).To(BeFalse())
	})
}

func TestUser_IsAccountComplete(t *testing.T) {
	t.Run("returns true when BranchOfService is set", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{BranchOfService: "Army"}

		g.Expect(u.IsAccountComplete()).To(BeTrue())
	})

	t.Run("returns false when BranchOfService is empty", func(t *testing.T) {
		g := NewWithT(t)

		u := &User{BranchOfService: ""}

		g.Expect(u.IsAccountComplete()).To(BeFalse())
	})
}
