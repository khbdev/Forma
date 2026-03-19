package domain

import (
	"context"
	"time"

	"forma/internal/models"
)

type LeadCache interface {
	GetAll(ctx context.Context) ([]models.Lead, error)
	SetAll(ctx context.Context, leads []models.Lead, ttl time.Duration) error
	DeleteAll(ctx context.Context) error

	GetByID(ctx context.Context, id uint) (*models.Lead, error)
	SetByID(ctx context.Context, lead *models.Lead, ttl time.Duration) error
	DeleteByID(ctx context.Context, id uint) error
}