package model

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestHunt_TotalCapacity(t *testing.T) {
	t.Run("returns sum of primary and alternate", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			PrimaryCapacity:   10,
			AlternateCapacity: 5,
		}

		g.Expect(h.TotalCapacity()).To(Equal(15))
	})

	t.Run("returns zero when both are zero", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			PrimaryCapacity:   0,
			AlternateCapacity: 0,
		}

		g.Expect(h.TotalCapacity()).To(Equal(0))
	})
}

func TestHunt_IsSignupWindowOpen(t *testing.T) {
	baseTime := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

	t.Run("returns true when within window", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			SignupWindowStart: baseTime.Add(-24 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.IsSignupWindowOpen(baseTime)).To(BeTrue())
	})

	t.Run("returns false when before window", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			SignupWindowStart: baseTime.Add(1 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.IsSignupWindowOpen(baseTime)).To(BeFalse())
	})

	t.Run("returns false when after window", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			SignupWindowStart: baseTime.Add(-48 * time.Hour),
			SignupWindowEnd:   baseTime.Add(-24 * time.Hour),
		}

		g.Expect(h.IsSignupWindowOpen(baseTime)).To(BeFalse())
	})
}

func TestHunt_CanAcceptSignups(t *testing.T) {
	baseTime := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

	t.Run("returns true when open and within window", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			Status:            HuntStatusOpen,
			SignupWindowStart: baseTime.Add(-24 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.CanAcceptSignups(baseTime)).To(BeTrue())
	})

	t.Run("returns false when draft", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			Status:            HuntStatusDraft,
			SignupWindowStart: baseTime.Add(-24 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.CanAcceptSignups(baseTime)).To(BeFalse())
	})

	t.Run("returns false when closed", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			Status:            HuntStatusClosed,
			SignupWindowStart: baseTime.Add(-24 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.CanAcceptSignups(baseTime)).To(BeFalse())
	})

	t.Run("returns false when open but outside window", func(t *testing.T) {
		g := NewWithT(t)

		h := &Hunt{
			Status:            HuntStatusOpen,
			SignupWindowStart: baseTime.Add(1 * time.Hour),
			SignupWindowEnd:   baseTime.Add(24 * time.Hour),
		}

		g.Expect(h.CanAcceptSignups(baseTime)).To(BeFalse())
	})
}
