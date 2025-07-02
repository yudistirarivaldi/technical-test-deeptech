package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
)

type CategoriesService struct {
	Repo *repository.CategoriesRepository
}

func NewCategoriesService(repo *repository.CategoriesRepository) *CategoriesService {
	return &CategoriesService{Repo: repo}
}

func (s *CategoriesService) InsertCategory(ctx context.Context, category *model.Categories) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}

	if err := s.Repo.InsertCategories(ctx, category); err != nil {
		return err
	}

	return nil
}

func (s *CategoriesService) GetAll(ctx context.Context) ([]*model.Categories, error) {
	categories, err := s.Repo.GetAllCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return categories, nil
}

func (s *CategoriesService) GetByID(ctx context.Context, id int64) (*model.Categories, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid category ID")
	}

	category, err := s.Repo.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %w", err)
	}
	return category, nil
}

func (s *CategoriesService) UpdateCategory(ctx context.Context, category *model.Categories) error {
	if category == nil {
		return fmt.Errorf("category is nil")
	}
	if category.ID == 0 {
		return fmt.Errorf("missing category ID")
	}

	err := s.Repo.UpdateCategory(ctx, category)
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}
	return nil
}

func (s *CategoriesService) DeleteCategory(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid category ID")
	}

	return s.Repo.DeleteCategory(ctx, id)
}
