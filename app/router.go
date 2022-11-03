package app

import (
	"project_spk_pemilihan_tabungan/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(adminController controllers.AdminController, tabunganController controllers.TabunganController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/register", adminController.AdminRegister)
			admin.POST("/login", adminController.AdminLogin)
			admin.POST("/logout", adminController.AdminLogout)
		}

		tabungan := api.Group("/tabungan")
		{
			tabungan.POST("/create", tabunganController.CreateTabungan)
			tabungan.GET("/list", tabunganController.ListAllTabungan)
			detail := tabungan.Group("detail/:id")
			{
				detail.GET("/", tabunganController.DetailTabungan)
				detail.PUT("/update", tabunganController.UpdateTabungan)
				detail.DELETE("/delete", tabunganController.DeleteTabungan)
			}
		}
	}

	return router
}
