package repositories

import (
	"context"
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type InputRecomendationRepository interface {
	Create(ctx context.Context, DB *gorm.DB, input *model.InputRecomendation) (*model.InputRecomendation, error)
}
