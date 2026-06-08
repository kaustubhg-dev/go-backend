package service

import (
	"context"
	"github.com/google/uuid"

	"go-backend/internal/models"

)

type ProductService interface {
	Create(ctx context.Context, req models.CreateProductRequest) (*models.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	GetAll(ctx context.Context, page, limit int) ([]models.Product, int64, error)
	Update(ctx context.Context, id uuid.UUID, req models.CreateProductRequest) (*models.Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}