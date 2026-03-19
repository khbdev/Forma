package domain

import (
	"context"
	"forma/internal/models"
)

type LeadRepository interface {
	Create(ctx context.Context, lead *models.Lead) error
	GetAll(ctx context.Context) ([]models.Lead, error)
	GetByID(ctx context.Context, id uint) (*models.Lead, error)
	Delete(ctx context.Context, id uint) error
}