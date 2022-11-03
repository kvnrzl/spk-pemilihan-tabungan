package repositories

import (
	"context"
	"project_spk_pemilihan_tabungan/models"
	"time"

	"gorm.io/gorm"
)

type TabunganRepositoryImpl struct{}

func NewTabunganRepository() TabunganRepository {
	return &TabunganRepositoryImpl{}
}

func (r *TabunganRepositoryImpl) Create(ctx context.Context, DB *gorm.DB, tabungan *models.Tabungan) (*models.Tabungan, error) {
	result := DB.WithContext(ctx).Create(&tabungan)
	if result.Error != nil {
		return &models.Tabungan{}, result.Error
	}

	tabungan.CreatedAt = time.Now()
	tabungan.UpdatedAt = time.Now()

	return tabungan, nil

	// return &models.Tabungan{
	// 	ID:                     tabungan.ID,
	// 	NamaTabungan:           tabungan.NamaTabungan,
	// 	SetoranAwal:            tabungan.SetoranAwal,
	// 	SetoranLanjutanMinimal: tabungan.SetoranLanjutanMinimal,
	// 	SaldoMinimum:           tabungan.SaldoMinimum,
	// 	SukuBunga:              tabungan.SukuBunga,
	// 	Fungsionalitas:         tabungan.Fungsionalitas,
	// 	BiayaAdmin:             tabungan.BiayaAdmin,
	// 	BiayaPenarikanHabis:    tabungan.BiayaPenarikanHabis,
	// 	KategoriUmurPengguna:   tabungan.KategoriUmurPengguna,
	// 	CreatedAt:              tabungan.CreatedAt,
	// 	UpdatedAt:              tabungan.UpdatedAt,
	// }, nil
}

func (r *TabunganRepositoryImpl) FindById(ctx context.Context, DB *gorm.DB, id int) (*models.Tabungan, error) {
	var tabungan models.Tabungan
	result := DB.WithContext(ctx).Where("id = ?", id).First(&tabungan)
	if result.Error != nil {
		return &models.Tabungan{}, result.Error
	}

	return &tabungan, nil
}

func (r *TabunganRepositoryImpl) FindAll(ctx context.Context, DB *gorm.DB) ([]models.Tabungan, error) {
	var tabungan []models.Tabungan
	result := DB.WithContext(ctx).Find(&tabungan)
	if result.Error != nil {
		return []models.Tabungan{}, result.Error
	}

	return tabungan, nil
}

func (r *TabunganRepositoryImpl) Update(ctx context.Context, DB *gorm.DB, tabungan *models.Tabungan) (*models.Tabungan, error) {
	result := DB.WithContext(ctx).Model(&tabungan).Where("id = ?", tabungan.ID).Updates(&tabungan)
	if result.Error != nil {
		return &models.Tabungan{}, result.Error
	}

	tabungan.UpdatedAt = time.Now()

	return tabungan, nil
}

func (r *TabunganRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, id int) error {
	result := DB.WithContext(ctx).Delete(&models.Tabungan{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
