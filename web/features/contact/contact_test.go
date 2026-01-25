package contact

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultRegionalLeaders(t *testing.T) {
	t.Run("returns placeholder regional leaders", func(t *testing.T) {
		g := NewWithT(t)

		leaders := DefaultRegionalLeaders()

		g.Expect(leaders).To(HaveLen(6))
		g.Expect(leaders[0].Name).ToNot(BeEmpty())
		g.Expect(leaders[0].Region).ToNot(BeEmpty())
		g.Expect(leaders[0].Email).ToNot(BeEmpty())
	})
}
