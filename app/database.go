package app

import (
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/project_spk_pemilihan_tabungan?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.Admin{}); err != nil {
		panic(err)
	}

	return DB
}
