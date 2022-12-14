package app

import (
	"project_spk_pemilihan_tabungan/config"
	model "project_spk_pemilihan_tabungan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")" + "/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.Admin{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.Tabungan{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.PresetKriteria{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.NilaiIdeal{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.Fungsionalitas{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.KategoriUmurPengguna{}); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&model.InputRecomendation{}); err != nil {
		panic(err)
	}

	// if err := DB.AutoMigrate(&model.TabunganSkor{}); err != nil {
	// 	panic(err)
	// }

	if err := DB.AutoMigrate(&model.BobotKriteria{}); err != nil {
		panic(err)
	}

	// if err := DB.AutoMigrate(&controllers.InputRecomendation{}); err != nil {
	// 	panic(err)
	// }
	return DB
}
