package services

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TabunganServiceImpl struct {
	tabunganRepository repositories.TabunganRepository
	validate           *validator.Validate
	DB                 *gorm.DB
}

func NewTabunganService(tabunganRepository repositories.TabunganRepository, validate *validator.Validate, DB *gorm.DB) TabunganService {
	return &TabunganServiceImpl{
		tabunganRepository: tabunganRepository,
		validate:           validate,
		DB:                 DB,
	}
}

func (s *TabunganServiceImpl) CreateTabungan(ctx context.Context, tabungan *models.Tabungan) (*models.Tabungan, error) {
	err := s.validate.Struct(tabungan)
	if err != nil {
		return &models.Tabungan{}, err
	}

	return s.tabunganRepository.Create(ctx, s.DB, tabungan)
}

func (s *TabunganServiceImpl) FindTabunganById(ctx context.Context, id int) (*models.Tabungan, error) {
	return s.tabunganRepository.FindById(ctx, s.DB, id)
}

func (s *TabunganServiceImpl) FindAllTabungan(ctx context.Context) ([]models.Tabungan, error) {
	return s.tabunganRepository.FindAll(ctx, s.DB)
}

func (s *TabunganServiceImpl) UpdateTabungan(ctx context.Context, tabungan *models.Tabungan) (*models.Tabungan, error) {
	err := s.validate.Struct(tabungan)
	if err != nil {
		return &models.Tabungan{}, err

	}

	// cek dulu datanya ada atau tidak
	_, err = s.tabunganRepository.FindById(ctx, s.DB, int(tabungan.ID))
	if err != nil {
		return &models.Tabungan{}, err
	}

	return s.tabunganRepository.Update(ctx, s.DB, tabungan)
}

func (s *TabunganServiceImpl) DeleteTabungan(ctx context.Context, id int) error {
	// cek dulu datanya ada atau tidak
	_, err := s.tabunganRepository.FindById(ctx, s.DB, id)
	if err != nil {
		return err
	}

	return s.tabunganRepository.Delete(ctx, s.DB, id)
}
