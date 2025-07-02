package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) InsertCategories(ctx context.Context, tx *model.Categories) error {
	query := `
		INSERT INTO categories (name, description)
		VALUES (?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, tx.Name, tx.Description)
	if err != nil {
		return fmt.Errorf("failed to insert category: %w", err)
	}
	return nil
}

func (r *CategoriesRepository) GetAllCategories(ctx context.Context) ([]*model.Categories, error) {
	query := `
		SELECT id, name, description
		FROM categories
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var results []*model.Categories
	for rows.Next() {
		var t model.Categories
		err := rows.Scan(&t.ID, &t.Name, &t.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		results = append(results, &t)
	}

	return results, nil
}

func (r *CategoriesRepository) GetCategoryByID(ctx context.Context, id int64) (*model.Categories, error) {
	query := `
		SELECT id, name, description
		FROM categories
		WHERE id = ?
	`

	var c model.Categories
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}
	return &c, nil
}

func (r *CategoriesRepository) UpdateCategory(ctx context.Context, category *model.Categories) error {
	query := `
		UPDATE categories
		SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Description, category.ID)
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}
	return nil
}

func (r *CategoriesRepository) DeleteCategory(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
