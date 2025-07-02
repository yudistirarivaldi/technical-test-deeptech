package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Insert(ctx context.Context, p *model.Product) error {
	return s.repo.Insert(ctx, p)
}

func (s *ProductService) GetAll(ctx context.Context) ([]*model.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id int64) (*model.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid product ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, p *model.Product) error {
	if p.ID <= 0 {
		return fmt.Errorf("invalid product ID")
	}
	return s.repo.Update(ctx, p)
}

func (s *ProductService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid product ID")
	}
	return s.repo.Delete(ctx, id)
}
