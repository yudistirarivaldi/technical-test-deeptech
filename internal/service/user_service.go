package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*model.Users, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid consumer ID")
	}

	consumer, err := s.Repo.GetByIDUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get consumer: %w", err)
	}
	if consumer == nil {
		return nil, nil
	}

	return consumer, nil
}

func (s *UserService) Update(ctx context.Context, consumer *model.Users) error {
	if consumer == nil {
		return fmt.Errorf("consumer is nil")
	}
	if consumer.ID == 0 {
		return fmt.Errorf("missing consumer ID")
	}

	err := s.Repo.UpdateConsumer(ctx, consumer)
	if err != nil {
		return fmt.Errorf("failed to update consumer: %w", err)
	}

	return nil
}
