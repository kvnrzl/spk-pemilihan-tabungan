package services

import (
	"context"
	"fmt"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PresetKriteriaServiceImpl struct {
	presetKriteriaRepository repositories.PresetKriteriaRepository
	validate                 *validator.Validate
	DB                       *gorm.DB
}

func NewPresetKriteriaService(presetKriteriaRepository repositories.PresetKriteriaRepository, validate *validator.Validate, DB *gorm.DB) PresetKriteriaService {
	return &PresetKriteriaServiceImpl{
		presetKriteriaRepository: presetKriteriaRepository,
		validate:                 validate,
		DB:                       DB,
	}
}

func (s *PresetKriteriaServiceImpl) CreatePreset(ctx context.Context, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error) {
	err := s.validate.Struct(presetKriteria)
	if err != nil {
		return &models.PresetKriteria{}, err
	}

	// kalo misalnya belum ada data preset, maka akan dicreate data preset baru
	_, err = s.presetKriteriaRepository.FindFirst(ctx, s.DB)
	if err != nil {
		presetKriteria.CreatedAt = time.Now()
		presetKriteria.UpdatedAt = time.Now()

		return s.presetKriteriaRepository.Create(ctx, s.DB, presetKriteria)
		// return &models.PresetKriteria{}, err
	}

	return &models.PresetKriteria{}, fmt.Errorf("Preset kriteria already exists")
}

func (s *PresetKriteriaServiceImpl) FindFirstPreset(ctx context.Context) (*models.PresetKriteria, error) {
	return s.presetKriteriaRepository.FindFirst(ctx, s.DB)
}

// func (s *PresetKriteriaServiceImpl) FindAllPreset(ctx context.Context) ([]models.PresetKriteria, error) {
// 	return s.presetKriteriaRepository.FindAll(ctx, s.DB)
// }

func (s *PresetKriteriaServiceImpl) UpdatePreset(ctx context.Context, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error) {
	err := s.validate.Struct(presetKriteria)
	if err != nil {
		return &models.PresetKriteria{}, err

	}

	// cek dulu datanya ada atau tidak, kalo tidak ada maka return err
	result, err := s.presetKriteriaRepository.FindFirst(ctx, s.DB)
	if err != nil {
		return &models.PresetKriteria{}, err
	}

	presetKriteria.ID = result.ID
	presetKriteria.UpdatedAt = time.Now()

	return s.presetKriteriaRepository.Update(ctx, s.DB, presetKriteria)
}
