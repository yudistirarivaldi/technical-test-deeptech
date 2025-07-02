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

func (r *TransactionRepository) GetTransactionsByUserID(ctx context.Context, userID int64) ([]model.TransactionWithItems, error) {
	query := `
	SELECT 
		t.id AS transaction_id,
		t.transaction_type,
		t.user_id,
		ti.id AS transaction_item_id,
		ti.product_id,
		ti.quantity
	FROM 
		transactions t
	LEFT JOIN 
		transaction_items ti ON t.id = ti.transaction_id
	WHERE 
		t.user_id = ?
	ORDER BY 
		t.id DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactionsMap := make(map[int64]*model.TransactionWithItems)
	for rows.Next() {
		var (
			tid       int64
			tType     string
			uid       int64
			tiid      sql.NullInt64
			productID sql.NullInt64
			quantity  sql.NullInt64
		)

		if err := rows.Scan(&tid, &tType, &uid, &tiid, &productID, &quantity); err != nil {
			return nil, err
		}

		if _, exists := transactionsMap[tid]; !exists {
			transactionsMap[tid] = &model.TransactionWithItems{
				ID:              tid,
				UserID:          uid,
				TransactionType: tType,
				Items:           []model.TransactionItem{},
			}
		}

		if tiid.Valid {
			transactionsMap[tid].Items = append(transactionsMap[tid].Items, model.TransactionItem{
				ID:            tiid.Int64,
				ProductID:     productID.Int64,
				Quantity:      quantity.Int64,
				TransactionID: tid,
			})
		}
	}

	var results []model.TransactionWithItems
	for _, v := range transactionsMap {
		results = append(results, *v)
	}

	return results, nil
}
