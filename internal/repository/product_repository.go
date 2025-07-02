package repository

import (
	"context"
	"database/sql"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Insert(ctx context.Context, p *model.Product) error {
	query := `INSERT INTO products (name, description, image_url, category_id, stock) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.ImageURL, p.CategoryID, p.Stock)
	return err
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, description, image_url, category_id, stock FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.ImageURL, &p.CategoryID, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	var p model.Product
	err := r.db.QueryRowContext(ctx, `SELECT id, name, description, image_url, category_id, stock FROM products WHERE id = ?`, id).
		Scan(&p.ID, &p.Name, &p.Description, &p.ImageURL, &p.CategoryID, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &p, err
}

func (r *ProductRepository) Update(ctx context.Context, p *model.Product) error {
	query := `
		UPDATE products
		SET name = ?, description = ?, image_url = ?, category_id = ?, stock = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Description, p.ImageURL, p.CategoryID, p.Stock, p.ID)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id = ?`, id)
	return err
}
