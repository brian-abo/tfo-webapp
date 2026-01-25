package components

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultTestimonials(t *testing.T) {
	t.Run("returns placeholder testimonials", func(t *testing.T) {
		g := NewWithT(t)

		testimonials := DefaultTestimonials()

		g.Expect(testimonials).To(HaveLen(3))
		g.Expect(testimonials[0].Name).ToNot(BeEmpty())
		g.Expect(testimonials[0].Quote).ToNot(BeEmpty())
		g.Expect(testimonials[0].Detail).ToNot(BeEmpty())
	})
}
