package model

import (
	"database/sql"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestSignup_IsWithdrawn(t *testing.T) {
	t.Run("returns false when WithdrawnAt is null", func(t *testing.T) {
		g := NewWithT(t)

		s := &Signup{WithdrawnAt: sql.NullTime{Valid: false}}

		g.Expect(s.IsWithdrawn()).To(BeFalse())
	})

	t.Run("returns true when WithdrawnAt is set", func(t *testing.T) {
		g := NewWithT(t)

		s := &Signup{WithdrawnAt: sql.NullTime{Time: time.Now(), Valid: true}}

		g.Expect(s.IsWithdrawn()).To(BeTrue())
	})
}
