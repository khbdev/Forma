package repository

import (
	"context"
	"errors"
	"forma/internal/domain"
	"forma/internal/models"

	"gorm.io/gorm"
)

type leadRepository struct {
	db *gorm.DB
}

func NewLeadRepository(db *gorm.DB) domain.LeadRepository {
	return  &leadRepository{db: db}
}



func (r *leadRepository) Create(ctx context.Context, lead *models.Lead) error {
	return r.db.WithContext(ctx).Create(lead).Error
}

func (r *leadRepository) GetAll(ctx context.Context) ([]models.Lead, error) {
	var leads []models.Lead

	err := r.db.WithContext(ctx).
		Order("id DESC").
		Find(&leads).Error

	return leads, err
}

func (r *leadRepository) GetByID(ctx context.Context, id uint) (*models.Lead, error) {
	var lead models.Lead

	err := r.db.WithContext(ctx).First(&lead, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &lead, nil
}

func (r *leadRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Lead{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}