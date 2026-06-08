package service

import (
	"context"

	"github.com/google/uuid"
	"go-backend/config"
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"go-backend/internal/utils"
)

// UserService contract.
type UserService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, req models.LoginRequest) (*utils.TokenPair, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.UserResponse, error)
	GetAll(ctx context.Context, page, limit int) ([]models.UserResponse, int64, error)
	Update(ctx context.Context, id uuid.UUID, req models.UpdateUserRequest) (*models.UserResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func (s *userService) Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, error) {
	return nil, nil
}

func (s *userService) Login(ctx context.Context, req models.LoginRequest) (*utils.TokenPair, error) {
	return nil, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*models.UserResponse, error) {
	return nil, nil
}

func (s *userService) GetAll(ctx context.Context, page, limit int) ([]models.UserResponse, int64, error) {
	return nil, 0, nil
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, req models.UpdateUserRequest) (*models.UserResponse, error) {
	return nil, nil
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	user, _ := s.repo.FindByID(ctx, id)
	if user == nil {
		return utils.NewAppError(404, "user not found")
	}
	return s.repo.Delete(ctx, id)
}

func NewUserService(repo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		repo: repo,
		cfg:  cfg,
	}
}