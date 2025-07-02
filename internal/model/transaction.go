package model

type Transaction struct {
	ID              int64  `json:"id"`
	TransactionType string `json:"transaction_type"`
	UserID          int64  `json:"user_id"`
}

type TransactionItem struct {
	ID            int64 `json:"id"`
	TransactionID int64 `json:"transaction_id"`
	ProductID     int64 `json:"product_id"`
	Quantity      int64 `json:"quantity"`
}
