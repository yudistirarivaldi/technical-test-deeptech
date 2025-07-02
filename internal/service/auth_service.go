package service

import (
	"context"
	"fmt"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/model"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo  *repository.AuthRepository
	jwtSecret string
}

func NewAuthService(authRepo *repository.AuthRepository, jwtSecret string) *AuthService {
	return &AuthService{
		authRepo:  authRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, users *model.Users) (int64, error) {
	existing, err := s.authRepo.FindByEmail(ctx, users.Email)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, fmt.Errorf("Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}
	users.Password = string(hashedPassword)

	consumerID, err := s.authRepo.RegisterConsumer(ctx, users)
	if err != nil {
		return 0, err
	}

	return consumerID, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	users, err := s.authRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if users == nil {
		return "", fmt.Errorf("E not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(users.ID, s.jwtSecret)

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
