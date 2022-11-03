package repositories

import (
	"context"
	"project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type TabunganRepository interface {
	Create(ctx context.Context, DB *gorm.DB, tabungan *models.Tabungan) (*models.Tabungan, error)
	FindById(ctx context.Context, DB *gorm.DB, id int) (*models.Tabungan, error)
	FindAll(ctx context.Context, DB *gorm.DB) ([]models.Tabungan, error)
	Update(ctx context.Context, DB *gorm.DB, tabungan *models.Tabungan) (*models.Tabungan, error)
	Delete(ctx context.Context, DB *gorm.DB, id int) error
}
