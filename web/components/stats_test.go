package components

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultStats(t *testing.T) {
	t.Run("returns placeholder statistics", func(t *testing.T) {
		g := NewWithT(t)

		stats := DefaultStats()

		g.Expect(stats).To(HaveLen(4))
		g.Expect(stats[0].Value).To(Equal("5,000+"))
		g.Expect(stats[0].Label).To(Equal("Veterans Served"))
	})
}
