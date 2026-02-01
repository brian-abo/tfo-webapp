package contact

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDefaultRegionalLeaders(t *testing.T) {
	t.Run("returns placeholder regional leaders", func(t *testing.T) {
		g := NewWithT(t)

		leaders := DefaultRegionalLeaders()

		g.Expect(leaders).To(HaveLen(4))
		g.Expect(leaders[0].Name).ToNot(BeEmpty())
		g.Expect(leaders[0].Region).ToNot(BeEmpty())
		g.Expect(leaders[0].Email).ToNot(BeEmpty())
	})
}

func TestUSRegions(t *testing.T) {
	t.Run("returns 4 regions with correct state counts", func(t *testing.T) {
		g := NewWithT(t)

		regions := USRegions()

		g.Expect(regions).To(HaveLen(4))

		byID := make(map[string]Region, len(regions))
		for _, r := range regions {
			byID[r.ID] = r
		}

		g.Expect(byID).To(HaveKey("west-coast"))
		g.Expect(byID).To(HaveKey("midwest"))
		g.Expect(byID).To(HaveKey("east-coast"))
		g.Expect(byID).To(HaveKey("southern"))

		g.Expect(byID["west-coast"].States).To(HaveLen(13))
		g.Expect(byID["midwest"].States).To(HaveLen(12))
		g.Expect(byID["east-coast"].States).To(HaveLen(16))
		g.Expect(byID["southern"].States).To(HaveLen(10))
	})

	t.Run("each region has a color", func(t *testing.T) {
		g := NewWithT(t)

		for _, r := range USRegions() {
			g.Expect(r.Color).ToNot(BeEmpty(), "region %s missing color", r.ID)
		}
	})

	t.Run("leader IDs match region IDs", func(t *testing.T) {
		g := NewWithT(t)

		regionIDs := make(map[string]bool)
		for _, r := range USRegions() {
			regionIDs[r.ID] = true
		}

		for _, l := range DefaultRegionalLeaders() {
			g.Expect(regionIDs).To(HaveKey(l.ID), "leader %s has no matching region", l.ID)
		}
	})
}
