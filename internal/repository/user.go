package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/brian-abo/tfo-webapp/internal/model"
)

// UserRepository handles persistence of user data.
type UserRepository struct {
	db DBTX
}

// NewUserRepository creates a UserRepository backed by the given DBTX.
func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID retrieves a user by their ID.
func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return r.scanUser(r.db.QueryRowContext(ctx,
		`SELECT id, email, name, phone, branch_of_service, role, membership_status,
		        facebook_id, created_at, updated_at, deleted_at
		 FROM users
		 WHERE id = $1 AND deleted_at IS NULL`,
		id,
	))
}

// FindByEmail retrieves a user by their email address.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return r.scanUser(r.db.QueryRowContext(ctx,
		`SELECT id, email, name, phone, branch_of_service, role, membership_status,
		        facebook_id, created_at, updated_at, deleted_at
		 FROM users
		 WHERE email = $1 AND deleted_at IS NULL`,
		email,
	))
}

// FindByFacebookID retrieves a user by their Facebook ID.
func (r *UserRepository) FindByFacebookID(ctx context.Context, facebookID string) (*model.User, error) {
	return r.scanUser(r.db.QueryRowContext(ctx,
		`SELECT id, email, name, phone, branch_of_service, role, membership_status,
		        facebook_id, created_at, updated_at, deleted_at
		 FROM users
		 WHERE facebook_id = $1 AND deleted_at IS NULL`,
		facebookID,
	))
}

// Create inserts a new user and returns it.
func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	var created model.User
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (email, name, phone, branch_of_service, role, membership_status, facebook_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 RETURNING id, email, name, phone, branch_of_service, role, membership_status,
		           facebook_id, created_at, updated_at, deleted_at`,
		user.Email, user.Name, user.Phone, user.BranchOfService,
		user.Role, user.MembershipStatus, user.FacebookID,
	).Scan(&created.ID, &created.Email, &created.Name, &created.Phone,
		&created.BranchOfService, &created.Role, &created.MembershipStatus,
		&created.FacebookID, &created.CreatedAt, &created.UpdatedAt, &created.DeletedAt)
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	return &created, nil
}

// Update modifies an existing user.
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE users SET
			email = $1, name = $2, phone = $3, branch_of_service = $4,
			role = $5, membership_status = $6, facebook_id = $7, updated_at = NOW()
		 WHERE id = $8 AND deleted_at IS NULL`,
		user.Email, user.Name, user.Phone, user.BranchOfService,
		user.Role, user.MembershipStatus, user.FacebookID, user.ID,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}
	return nil
}

func (r *UserRepository) scanUser(row *sql.Row) (*model.User, error) {
	var u model.User
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Phone, &u.BranchOfService,
		&u.Role, &u.MembershipStatus, &u.FacebookID, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("scanning user: %w", err)
	}
	return &u, nil
}
