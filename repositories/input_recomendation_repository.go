package repositories

import (
	"context"
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type InputRecomendationRepositoryImpl struct{}

func NewInputRecomendationRepository() InputRecomendationRepository {
	return &InputRecomendationRepositoryImpl{}
}

func (r *InputRecomendationRepositoryImpl) Create(ctx context.Context, DB *gorm.DB, input *model.InputRecomendation) (*model.InputRecomendation, error) {
	err := DB.Create(input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}
