package services

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type InputRecomendationServiceImpl struct {
	inputRecomendationRepository repositories.InputRecomendationRepository
	db                           *gorm.DB
	validate                     *validator.Validate
}

func NewInputRecomendationService(inputRecomendationRepository repositories.InputRecomendationRepository, db *gorm.DB, validate *validator.Validate) InputRecomendationService {
	return &InputRecomendationServiceImpl{
		inputRecomendationRepository: inputRecomendationRepository,
		db:                           db,
		validate:                     validate,
	}
}

func (s *InputRecomendationServiceImpl) Create(ctx context.Context, input *models.InputRecomendation) (*models.InputRecomendation, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	result, err := s.inputRecomendationRepository.Create(ctx, s.db, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
