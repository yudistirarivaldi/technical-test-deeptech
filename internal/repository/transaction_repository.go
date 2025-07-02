package repository

import (
	"context"
	"database/sql"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *TransactionRepository) InsertTransaction(ctx context.Context, tx *sql.Tx, t *model.Transaction) (int64, error) {
	query := `INSERT INTO transactions (transaction_type, user_id) VALUES (?, ?)`
	res, err := tx.ExecContext(ctx, query, t.TransactionType, t.UserID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *TransactionRepository) InsertTransactionItem(ctx context.Context, tx *sql.Tx, item *model.TransactionItem) error {
	query := `INSERT INTO transaction_items (transaction_id, product_id, quantity) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, item.TransactionID, item.ProductID, item.Quantity)
	return err
}

func (r *TransactionRepository) GetProductStockForUpdate(ctx context.Context, tx *sql.Tx, productID int64) (int64, error) {
	var stock int64
	query := `SELECT stock FROM products WHERE id = ? FOR UPDATE`
	err := tx.QueryRowContext(ctx, query, productID).Scan(&stock)
	return stock, err
}

func (r *TransactionRepository) UpdateProductStock(ctx context.Context, tx *sql.Tx, productID, newStock int64) error {
	query := `UPDATE products SET stock = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, query, newStock, productID)
	return err
}
