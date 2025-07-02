package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) RegisterConsumer(ctx context.Context, c *model.Users) (int64, error) {
	query := `
		INSERT INTO users (
			first_name, last_name, email, password, date_of_birth, gender
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(ctx, query,
		c.FirstName,
		c.LastName,
		c.Email,
		c.Password,
		c.DateOfBirth,
		c.Gender,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to register users: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return insertedID, nil
}

func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (*model.Users, error) {
	query := `
		SELECT id, first_name, last_name, email, password, date_of_birth, gender
		FROM users WHERE email = ? LIMIT 1
	`

	var c model.Users
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&c.ID,
		&c.FirstName,
		&c.LastName,
		&c.Email,
		&c.Password,
		&c.DateOfBirth,
		&c.Gender,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find consumer by email: %w", err)
	}

	return &c, nil
}
