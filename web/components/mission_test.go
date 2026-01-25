package components

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultMissionProps(t *testing.T) {
	t.Run("returns standard mission content", func(t *testing.T) {
		g := NewWithT(t)

		props := DefaultMissionProps()

		g.Expect(props.Heading).To(Equal("Our Mission"))
		g.Expect(props.Description).ToNot(BeEmpty())
	})
}
