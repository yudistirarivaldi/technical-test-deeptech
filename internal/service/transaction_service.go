package service

import (
	"context"
	"fmt"
	"log"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/dto"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(ctx context.Context, req *dto.CreateTransactionRequest) error {
	log.Printf("[TransactionService] Creating transaction for user_id: %d, type: %s", req.UserID, req.TransactionType)

	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		log.Printf("[TransactionService] Failed to begin transaction: %v", err)
		return err
	}
	defer tx.Rollback()

	transaction := &model.Transaction{
		UserID:          req.UserID,
		TransactionType: req.TransactionType,
	}
	transactionID, err := s.repo.InsertTransaction(ctx, tx, transaction)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	for _, item := range req.Items {

		stock, err := s.repo.GetProductStockForUpdate(ctx, tx, item.ProductID)
		if err != nil {
			return fmt.Errorf("product not found or locked: %w", err)
		}

		newStock := stock
		if req.TransactionType == "IN" {
			newStock += item.Quantity
		} else if req.TransactionType == "OUT" {
			if item.Quantity > stock {

				return fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
			}
			newStock -= item.Quantity
		}

		if err := s.repo.UpdateProductStock(ctx, tx, item.ProductID, newStock); err != nil {

			return fmt.Errorf("failed to update stock: %w", err)
		}

		itemModel := &model.TransactionItem{
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
		}

		if err := s.repo.InsertTransactionItem(ctx, tx, itemModel); err != nil {

			return fmt.Errorf("failed to insert transaction item: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {

		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *TransactionService) GetByUserID(ctx context.Context, userID int64) ([]model.TransactionWithItems, error) {
	return s.repo.GetTransactionsByUserID(ctx, userID)
}
