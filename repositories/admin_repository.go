package repositories

import (
	"context"
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(ctx context.Context, DB *gorm.DB, admin *model.Admin) (*model.Admin, error)
	FindByUsername(ctx context.Context, DB *gorm.DB, username string) (*model.Admin, error)
	Update(ctx context.Context, DB *gorm.DB, admin *model.Admin) (*model.Admin, error)
	Delete(ctx context.Context, DB *gorm.DB, id int) error
}
