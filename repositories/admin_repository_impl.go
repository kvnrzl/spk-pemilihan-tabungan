package repositories

import (
	"context"
	"fmt"
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct{}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (r *AdminRepositoryImpl) Create(ctx context.Context, DB *gorm.DB, admin *model.Admin) (*model.Admin, error) {
	result := DB.WithContext(ctx).Create(&admin)
	if result.Error != nil {
		return &model.Admin{}, result.Error
	}

	return &model.Admin{
		Username: admin.Username,
	}, nil
}

func (r *AdminRepositoryImpl) FindByUsername(ctx context.Context, DB *gorm.DB, username string) (*model.Admin, error) {
	var admin model.Admin

	result := DB.WithContext(ctx).Where("username = ?", username).Find(&admin)
	if result.Error != nil {
		return &model.Admin{}, result.Error
	}

	if result.RowsAffected == 0 {
		return &model.Admin{}, fmt.Errorf("Username or password is wrong")
	}

	return &model.Admin{
		Username: admin.Username,
		Password: admin.Password,
	}, nil
}

func (r *AdminRepositoryImpl) Update(ctx context.Context, DB *gorm.DB, admin *model.Admin) (*model.Admin, error) {
	result := DB.WithContext(ctx).Where("id = ?", admin.ID).Updates(&admin)
	if result.Error != nil {
		return &model.Admin{}, result.Error
	}

	return &model.Admin{
		Username: admin.Username,
	}, nil
}

func (r *AdminRepositoryImpl) Delete(ctx context.Context, DB *gorm.DB, id int) error {
	result := DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Admin{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
