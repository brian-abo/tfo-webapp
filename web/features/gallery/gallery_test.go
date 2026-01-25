package gallery

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultGalleryImages(t *testing.T) {
	t.Run("returns placeholder images", func(t *testing.T) {
		g := NewWithT(t)

		images := DefaultGalleryImages()

		g.Expect(images).To(HaveLen(12))
		g.Expect(images[0].Src).ToNot(BeEmpty())
		g.Expect(images[0].Alt).ToNot(BeEmpty())
		g.Expect(images[0].Caption).ToNot(BeEmpty())
	})
}
