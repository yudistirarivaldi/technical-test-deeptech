package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByIDUser(ctx context.Context, id int64) (*model.Users, error) {
	query := `
		SELECT id, first_name, last_name, email, date_of_birth, gender
		FROM users WHERE id = ?
	`

	var c model.Users
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.Email,
		&c.DateOfBirth,
		&c.Gender,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get consumer by ID: %w", err)
	}

	return &c, nil
}

func (r *UserRepository) UpdateConsumer(ctx context.Context, c *model.Users) error {
	query := `
		UPDATE users
		SET first_name = ?, last_name = ?, email = ?, password = ?, date_of_birth = ?, 
		    gender = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query,
		c.FirstName,
		c.LastName,
		c.Email,
		c.Password,
		c.DateOfBirth,
		c.Gender,
		c.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update consumer: %w", err)
	}

	return nil
}
