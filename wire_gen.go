// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"project_spk_pemilihan_tabungan/app"
	"project_spk_pemilihan_tabungan/controllers"
	"project_spk_pemilihan_tabungan/repositories"
	"project_spk_pemilihan_tabungan/services"
)

// Injectors from injector.go:

func InitServer() *gin.Engine {
	adminRepository := repositories.NewAdminRepository()
	db := app.InitDB()
	validate := validator.New()
	adminService := services.NewAdminService(adminRepository, db, validate)
	adminController := controllers.NewAdminController(adminService)
	engine := app.NewRouter(adminController)
	return engine
}
