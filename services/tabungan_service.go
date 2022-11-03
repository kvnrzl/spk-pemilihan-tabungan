package services

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
)

type TabunganService interface {
	CreateTabungan(ctx context.Context, tabungan *models.Tabungan) (*models.Tabungan, error)
	FindTabunganById(ctx context.Context, id int) (*models.Tabungan, error)
	FindAllTabungan(ctx context.Context) ([]models.Tabungan, error)
	UpdateTabungan(ctx context.Context, tabungan *models.Tabungan) (*models.Tabungan, error)
	DeleteTabungan(ctx context.Context, id int) error
}
