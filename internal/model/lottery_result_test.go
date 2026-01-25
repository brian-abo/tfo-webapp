package model

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestLotteryResult_IsPrimary(t *testing.T) {
	t.Run("returns true when position is within primary capacity", func(t *testing.T) {
		g := NewWithT(t)

		r := &LotteryResult{Position: 5}

		g.Expect(r.IsPrimary(10)).To(BeTrue())
	})

	t.Run("returns true when position equals primary capacity", func(t *testing.T) {
		g := NewWithT(t)

		r := &LotteryResult{Position: 10}

		g.Expect(r.IsPrimary(10)).To(BeTrue())
	})

	t.Run("returns false when position exceeds primary capacity", func(t *testing.T) {
		g := NewWithT(t)

		r := &LotteryResult{Position: 11}

		g.Expect(r.IsPrimary(10)).To(BeFalse())
	})

	t.Run("returns true for position 1 with any capacity", func(t *testing.T) {
		g := NewWithT(t)

		r := &LotteryResult{Position: 1}

		g.Expect(r.IsPrimary(1)).To(BeTrue())
	})
}
