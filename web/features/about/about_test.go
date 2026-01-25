package about

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultLeaders(t *testing.T) {
	t.Run("returns placeholder leaders", func(t *testing.T) {
		g := NewWithT(t)

		leaders := DefaultLeaders()

		g.Expect(leaders).To(HaveLen(4))
		g.Expect(leaders[0].Name).ToNot(BeEmpty())
		g.Expect(leaders[0].Title).ToNot(BeEmpty())
		g.Expect(leaders[0].Bio).ToNot(BeEmpty())
	})
}

func TestDefaultDocuments(t *testing.T) {
	t.Run("returns placeholder documents", func(t *testing.T) {
		g := NewWithT(t)

		docs := DefaultDocuments()

		g.Expect(docs).To(HaveLen(4))
		g.Expect(docs[0].Name).ToNot(BeEmpty())
		g.Expect(docs[0].Description).ToNot(BeEmpty())
		g.Expect(docs[0].Href).ToNot(BeEmpty())
	})
}
