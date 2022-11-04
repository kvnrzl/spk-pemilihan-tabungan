package repositories

import (
	"context"
	"project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type PresetKriteriaRepository interface {
	Create(ctx context.Context, DB *gorm.DB, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error)
	FindFirst(ctx context.Context, DB *gorm.DB) (*models.PresetKriteria, error)
	// FindAll(ctx context.Context, DB *gorm.DB) ([]models.PresetKriteria, error)
	Update(ctx context.Context, DB *gorm.DB, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error)
}
