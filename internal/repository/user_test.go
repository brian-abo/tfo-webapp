package repository_test

import (
	"database/sql"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/brian-abo/tfo-webapp/internal/model"
	"github.com/brian-abo/tfo-webapp/internal/repository"
)

func TestUserRepository_Create(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewUserRepository(tx)

		user := &model.User{
			Email:            "test@example.com",
			Name:             "Test User",
			BranchOfService:  "Army",
			Role:             model.RoleMember,
			MembershipStatus: model.MembershipPending,
			FacebookID:       sql.NullString{String: "fb123", Valid: true},
		}

		created, err := repo.Create(t.Context(), user)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(created.ID.String()).ToNot(BeEmpty())
		g.Expect(created.Email).To(Equal("test@example.com"))
		g.Expect(created.Name).To(Equal("Test User"))
		g.Expect(created.BranchOfService).To(Equal("Army"))
		g.Expect(created.Role).To(Equal(model.RoleMember))
		g.Expect(created.MembershipStatus).To(Equal(model.MembershipPending))
		g.Expect(created.FacebookID.String).To(Equal("fb123"))
		g.Expect(created.CreatedAt.IsZero()).To(BeFalse())
	})
}

func TestUserRepository_FindByID(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewUserRepository(tx)

		user := &model.User{
			Email:            "find@example.com",
			Name:             "Find User",
			BranchOfService:  "Navy",
			Role:             model.RoleMember,
			MembershipStatus: model.MembershipActive,
		}
		created, err := repo.Create(t.Context(), user)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := repo.FindByID(t.Context(), created.ID)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).ToNot(BeNil())
		g.Expect(found.Email).To(Equal("find@example.com"))
	})
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewUserRepository(tx)

		user := &model.User{
			Email:            "email@example.com",
			Name:             "Email User",
			BranchOfService:  "Marines",
			Role:             model.RoleMember,
			MembershipStatus: model.MembershipActive,
		}
		_, err := repo.Create(t.Context(), user)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := repo.FindByEmail(t.Context(), "email@example.com")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).ToNot(BeNil())
		g.Expect(found.Name).To(Equal("Email User"))

		notFound, err := repo.FindByEmail(t.Context(), "nonexistent@example.com")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(notFound).To(BeNil())
	})
}

func TestUserRepository_FindByFacebookID(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewUserRepository(tx)

		user := &model.User{
			Email:            "fb@example.com",
			Name:             "Facebook User",
			BranchOfService:  "Air Force",
			Role:             model.RoleMember,
			MembershipStatus: model.MembershipActive,
			FacebookID:       sql.NullString{String: "fb456", Valid: true},
		}
		_, err := repo.Create(t.Context(), user)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := repo.FindByFacebookID(t.Context(), "fb456")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found).ToNot(BeNil())
		g.Expect(found.Name).To(Equal("Facebook User"))

		notFound, err := repo.FindByFacebookID(t.Context(), "nonexistent")
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(notFound).To(BeNil())
	})
}

func TestUserRepository_Update(t *testing.T) {
	db := testDB(t)

	withTestTx(t, db, func(tx *sql.Tx) {
		g := NewWithT(t)
		repo := repository.NewUserRepository(tx)

		user := &model.User{
			Email:            "update@example.com",
			Name:             "Original Name",
			BranchOfService:  "Coast Guard",
			Role:             model.RoleMember,
			MembershipStatus: model.MembershipPending,
		}
		created, err := repo.Create(t.Context(), user)
		g.Expect(err).ToNot(HaveOccurred())

		created.Name = "Updated Name"
		created.MembershipStatus = model.MembershipActive
		err = repo.Update(t.Context(), created)
		g.Expect(err).ToNot(HaveOccurred())

		found, err := repo.FindByID(t.Context(), created.ID)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(found.Name).To(Equal("Updated Name"))
		g.Expect(found.MembershipStatus).To(Equal(model.MembershipActive))
	})
}
