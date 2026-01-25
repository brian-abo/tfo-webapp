package components

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestCurrentYear(t *testing.T) {
	t.Run("returns current year as string", func(t *testing.T) {
		g := NewWithT(t)

		result := CurrentYear()
		expected := time.Now().Format("2006")

		g.Expect(result).To(Equal(expected))
	})
}
