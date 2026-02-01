package repository_test

import (
	"database/sql"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/brian-abo/tfo-webapp/internal/repository"
)

func TestContactRepository_Insert(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewContactRepository(tx)

		sub, err := repo.Insert(t.Context(), "Jane Doe", "jane@example.com", "Hello there")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(sub.ID.String()).ToNot(BeEmpty())
		g.Expect(sub.Name).To(Equal("Jane Doe"))
		g.Expect(sub.Email).To(Equal("jane@example.com"))
		g.Expect(sub.Message).To(Equal("Hello there"))
		g.Expect(sub.CreatedAt.IsZero()).To(BeFalse())
	})
}

func TestContactRepository_List(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewContactRepository(tx)

		for _, name := range []string{"Alice", "Bob", "Charlie"} {
			_, err := repo.Insert(t.Context(), name, name+"@example.com", "Message "+name)
			g.Expect(err).ToNot(HaveOccurred())
		}

		t.Run("respects limit", func(t *testing.T) {
			g := NewWithT(t)
			subs, err := repo.List(t.Context(), 2)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(subs).To(HaveLen(2))
		})

		t.Run("returns all when limit is high", func(t *testing.T) {
			g := NewWithT(t)
			all, err := repo.List(t.Context(), 100)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(all).To(HaveLen(3))

			names := make([]string, len(all))
			for i, s := range all {
				names[i] = s.Name
			}
			g.Expect(names).To(ContainElements("Alice", "Bob", "Charlie"))
		})
	})
}
