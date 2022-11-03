//go:build wireinject
// +build wireinject

package main

import (
	"project_spk_pemilihan_tabungan/app"
	"project_spk_pemilihan_tabungan/controllers"
	"project_spk_pemilihan_tabungan/repositories"
	"project_spk_pemilihan_tabungan/services"
	"github.com/gin-gonic/gin"

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
		app.NewRouter,
	)
	return nil
}
