package controllers

import (
	"net/http"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"

	"github.com/gin-gonic/gin"
)

type PresetKriteriaControllerImpl struct {
	PresetKriteriaService services.PresetKriteriaService
}

func NewPresetKriteriaController(presetKriteriaService services.PresetKriteriaService) PresetKriteriaController {
	return &PresetKriteriaControllerImpl{
		PresetKriteriaService: presetKriteriaService,
	}
}

func (c *PresetKriteriaControllerImpl) CreatePreset(ctx *gin.Context) {
	if _, err := ctx.Cookie("jwt"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "Unauthorized",
		})
		return
	}

	var presetKriteria models.PresetKriteria
	if err := ctx.ShouldBindJSON(&presetKriteria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	result, err := c.PresetKriteriaService.CreatePreset(ctx, &presetKriteria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success create preset kriteria",
		"data":    result,
	})
}

func (c *PresetKriteriaControllerImpl) FindFirstPreset(ctx *gin.Context) {
	result, err := c.PresetKriteriaService.FindFirstPreset(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success get preset kriteria",
		"data":    result,
	})
}

func (c *PresetKriteriaControllerImpl) UpdatePreset(ctx *gin.Context) {
	if _, err := ctx.Cookie("jwt"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "Unauthorized",
		})
		return
	}

	var presetKriteria models.PresetKriteria
	if err := ctx.ShouldBindJSON(&presetKriteria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	result, err := c.PresetKriteriaService.UpdatePreset(ctx, &presetKriteria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success update preset kriteria",
		"data":    result,
	})
}
