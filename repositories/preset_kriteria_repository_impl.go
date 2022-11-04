package repositories

import (
	"context"
	"project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type PresetKriteriaRepositoryImpl struct{}

func NewPresetKriteriaRepository() PresetKriteriaRepository {
	return &PresetKriteriaRepositoryImpl{}
}

func (r *PresetKriteriaRepositoryImpl) Create(ctx context.Context, DB *gorm.DB, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error) {
	result := DB.WithContext(ctx).Create(&presetKriteria)
	if result.Error != nil {
		return &models.PresetKriteria{}, result.Error
	}

	return presetKriteria, nil
}

func (r *PresetKriteriaRepositoryImpl) FindFirst(ctx context.Context, DB *gorm.DB) (*models.PresetKriteria, error) {
	var presetKriteria models.PresetKriteria
	result := DB.WithContext(ctx).First(&presetKriteria)
	if result.Error != nil {
		return &models.PresetKriteria{}, result.Error
	}

	return &presetKriteria, nil
}

// func (r *PresetKriteriaRepositoryImpl) FindAll(ctx context.Context, DB *gorm.DB) ([]models.PresetKriteria, error) {
// 	var presetKriteria []models.PresetKriteria
// 	result := DB.WithContext(ctx).Find(&presetKriteria)
// 	if result.Error != nil {
// 		return []models.PresetKriteria{}, result.Error
// 	}

// 	return presetKriteria, nil
// }

func (r *PresetKriteriaRepositoryImpl) Update(ctx context.Context, DB *gorm.DB, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error) {
	result := DB.WithContext(ctx).Where("id = ?", presetKriteria.ID).Updates(&presetKriteria)
	if result.Error != nil {
		return &models.PresetKriteria{}, result.Error
	}

	return presetKriteria, nil
}
