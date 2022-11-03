package app

import (
	"project_spk_pemilihan_tabungan/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(adminController controllers.AdminController) *gin.Engine {
	r := gin.Default()

	r.POST("/register", adminController.AdminRegister)
	r.POST("/login", adminController.AdminLogin)
	r.POST("/logout", adminController.AdminLogout)

	return r
}
