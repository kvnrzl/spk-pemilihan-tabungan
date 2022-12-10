package app

import (
	"project_spk_pemilihan_tabungan/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(adminController controllers.AdminController, tabunganController controllers.TabunganController, presetController controllers.PresetKriteriaController, resultController controllers.ResultController) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())
	// router.RedirectTrailingSlash = false

	router.GET("/", func(r *gin.Context) {
		r.JSON(200, "API service is ready!")
	})

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.POST("/register", adminController.AdminRegister)
			admin.POST("/login", adminController.AdminLogin)
			admin.POST("/logout", VerifyToken(), adminController.AdminLogout)
		}

		tabungan := api.Group("/tabungan")
		{
			tabungan.POST("/create", VerifyToken(), tabunganController.CreateTabungan)
			tabungan.GET("/list", tabunganController.ListAllTabungan)
			detail := tabungan.Group("detail/:id")
			{
				detail.GET("/", tabunganController.DetailTabungan)
				detail.PUT("/update", VerifyToken(), tabunganController.UpdateTabungan)
				detail.DELETE("/delete", VerifyToken(), tabunganController.DeleteTabungan)
			}
		}

		presetKriteria := api.Group("/preset")
		{
			presetKriteria.GET("/", presetController.FindFirstPreset)
			presetKriteria.POST("/create", VerifyToken(), presetController.CreatePreset)
			presetKriteria.PUT("/update", VerifyToken(), presetController.UpdatePreset)
		}

		result := api.Group("/result")
		{
			result.POST("/hitung", resultController.HitungResult)
		}
	}

	return router
}
