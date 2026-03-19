package service

import (
	"context"

	"forma/internal/domain"
	"forma/internal/models"
)

type leadService struct {
	repo domain.LeadRepository
}

func NewLeadService(repo domain.LeadRepository) domain.LeadService {
	return &leadService{repo: repo}
}

func (s *leadService) Create(ctx context.Context, lead *models.Lead) error {
	return s.repo.Create(ctx, lead)
}

func (s *leadService) GetAll(ctx context.Context) ([]models.Lead, error) {
	return s.repo.GetAll(ctx)
}

func (s *leadService) GetByID(ctx context.Context, id uint) (*models.Lead, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *leadService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}