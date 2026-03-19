package service

import (
	"context"
	"time"

	"forma/internal/domain"
	"forma/internal/models"
)

type leadService struct {
	repo  domain.LeadRepository
	cache domain.LeadCache
	ttl   time.Duration
}

func NewLeadService(repo domain.LeadRepository, cache domain.LeadCache) domain.LeadService {
	return &leadService{
		repo:  repo,
		cache: cache,
		ttl:   5 * time.Minute,
	}
}

func (s *leadService) Create(ctx context.Context, lead *models.Lead) error {
	if err := s.repo.Create(ctx, lead); err != nil {
		return err
	}

	// yangi lead qo‘shilgani uchun list cache eskiradi
	_ = s.cache.DeleteAll(ctx)

	// agar ID to‘lib qaytsa, item cachega ham yozib qo‘yamiz
	_ = s.cache.SetByID(ctx, lead, s.ttl)

	return nil
}

func (s *leadService) GetAll(ctx context.Context) ([]models.Lead, error) {
	leads, err := s.cache.GetAll(ctx)
	if err != nil {
		// cache ishlamasa ham DBdan davom etamiz
		leads, err = s.repo.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return leads, nil
	}

	if leads != nil {
		return leads, nil
	}

	leads, err = s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetAll(ctx, leads, s.ttl)

	return leads, nil
}

func (s *leadService) GetByID(ctx context.Context, id uint) (*models.Lead, error) {
	lead, err := s.cache.GetByID(ctx, id)
	if err != nil {
		lead, err = s.repo.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return lead, nil
	}

	if lead != nil {
		return lead, nil
	}

	lead, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetByID(ctx, lead, s.ttl)

	return lead, nil
}

func (s *leadService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	// o‘chirilgandan keyin cache invalidation
	_ = s.cache.DeleteByID(ctx, id)
	_ = s.cache.DeleteAll(ctx)

	return nil
}