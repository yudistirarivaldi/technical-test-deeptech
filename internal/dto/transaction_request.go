package dto

type TransactionItemRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required,gt=0"`
}

type CreateTransactionRequest struct {
	TransactionType string                   `json:"transaction_type" validate:"required,oneof=IN OUT"`
	UserID          int64                    `json:"-"`
	Items           []TransactionItemRequest `json:"items" validate:"required,dive"`
}
