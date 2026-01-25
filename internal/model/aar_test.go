package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/gomega"
)

func TestHuntAfterActionReport_Fields(t *testing.T) {
	t.Run("struct holds all required fields", func(t *testing.T) {
		g := NewWithT(t)

		huntID := uuid.New()
		createdByID := uuid.New()
		now := time.Now()

		aar := &HuntAfterActionReport{
			ID:          uuid.New(),
			HuntID:      huntID,
			Description: "Great hunt with excellent participation.",
			ImageURLs:   []string{"https://example.com/img1.jpg", "https://example.com/img2.jpg"},
			CreatedByID: createdByID,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		g.Expect(aar.HuntID).To(Equal(huntID))
		g.Expect(aar.Description).To(Equal("Great hunt with excellent participation."))
		g.Expect(aar.ImageURLs).To(HaveLen(2))
		g.Expect(aar.CreatedByID).To(Equal(createdByID))
	})
}

func TestAARParticipant_Fields(t *testing.T) {
	t.Run("struct holds AAR and user references", func(t *testing.T) {
		g := NewWithT(t)

		aarID := uuid.New()
		userID := uuid.New()

		p := &AARParticipant{
			AARID:  aarID,
			UserID: userID,
		}

		g.Expect(p.AARID).To(Equal(aarID))
		g.Expect(p.UserID).To(Equal(userID))
	})
}
