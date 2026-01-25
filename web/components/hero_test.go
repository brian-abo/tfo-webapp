package components

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultHeroProps(t *testing.T) {
	t.Run("returns standard hero content", func(t *testing.T) {
		g := NewWithT(t)

		props := DefaultHeroProps()

		g.Expect(props.Headline).To(Equal("Adventure Awaits"))
		g.Expect(props.Subheadline).ToNot(BeEmpty())
		g.Expect(props.CTAText).To(Equal("Join a Hunt"))
		g.Expect(props.CTAHref).To(Equal("/hunts"))
	})
}
