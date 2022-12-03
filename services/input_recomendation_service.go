package services

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
)

type InputRecomendationService interface {
	Create(ctx context.Context, input *models.InputRecomendation) (*models.InputRecomendation, error)
}
