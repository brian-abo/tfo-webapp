package repository

import (
	"context"
	"fmt"

	"github.com/brian-abo/tfo-webapp/internal/model"
)

// ContactRepository handles persistence of contact form submissions.
type ContactRepository struct {
	db DBTX
}

// NewContactRepository creates a ContactRepository backed by the given DBTX.
func NewContactRepository(db DBTX) *ContactRepository {
	return &ContactRepository{db: db}
}

// Insert stores a new contact form submission.
func (r *ContactRepository) Insert(ctx context.Context, name, email, message string) (model.ContactSubmission, error) {
	var s model.ContactSubmission
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO contact_submissions (name, email, message)
		 VALUES ($1, $2, $3)
		 RETURNING id, name, email, message, created_at`,
		name, email, message,
	).Scan(&s.ID, &s.Name, &s.Email, &s.Message, &s.CreatedAt)
	if err != nil {
		return model.ContactSubmission{}, fmt.Errorf("inserting contact submission: %w", err)
	}
	return s, nil
}

// List returns the most recent contact submissions, up to limit.
func (r *ContactRepository) List(ctx context.Context, limit int) ([]model.ContactSubmission, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, name, email, message, created_at
		 FROM contact_submissions
		 ORDER BY created_at DESC
		 LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("listing contact submissions: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var submissions []model.ContactSubmission
	for rows.Next() {
		var s model.ContactSubmission
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Message, &s.CreatedAt); err != nil {
			return nil, fmt.Errorf("scanning contact submission: %w", err)
		}
		submissions = append(submissions, s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating contact submissions: %w", err)
	}
	return submissions, nil
}
