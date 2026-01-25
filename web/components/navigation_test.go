package components

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultNavItems(t *testing.T) {
	t.Run("returns standard navigation items", func(t *testing.T) {
		g := NewWithT(t)

		items := DefaultNavItems()

		g.Expect(items).To(HaveLen(4))
		g.Expect(items[0].Label).To(Equal("Home"))
		g.Expect(items[0].Href).To(Equal("/"))
		g.Expect(items[1].Label).To(Equal("About"))
		g.Expect(items[3].Label).To(Equal("Contact"))
	})
}
