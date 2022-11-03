//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"project_spk_pemilihan_tabungan/app"
	"project_spk_pemilihan_tabungan/controllers"
	"project_spk_pemilihan_tabungan/repositories"
	"project_spk_pemilihan_tabungan/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitServer() *gin.Engine {
	wire.Build(
		app.InitDB,
		validator.New,
		repositories.NewAdminRepository,
		services.NewAdminService,
		controllers.NewAdminController,

		repositories.NewTabunganRepository,
		services.NewTabunganService,
		controllers.NewTabunganController,

		app.NewRouter,
	)
	return nil
}
