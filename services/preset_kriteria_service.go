package services

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
)

type PresetKriteriaService interface {
	CreatePreset(ctx context.Context, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error)
	FindFirstPreset(ctx context.Context) (*models.PresetKriteria, error)
	// FindAllPreset(ctx context.Context) ([]models.PresetKriteria, error)
	UpdatePreset(ctx context.Context, presetKriteria *models.PresetKriteria) (*models.PresetKriteria, error)
}
