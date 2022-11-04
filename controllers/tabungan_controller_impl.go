package controllers

import (
	"net/http"
	"project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TabunganControllerImpl struct {
	tabunganService services.TabunganService
}

func NewTabunganController(tabunganService services.TabunganService) TabunganController {
	return &TabunganControllerImpl{
		tabunganService: tabunganService,
	}
}

func (c *TabunganControllerImpl) CreateTabungan(ctx *gin.Context) {
	if _, err := ctx.Cookie("jwt"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": err.Error(),
		})
		return
	}

	var tabungan models.Tabungan
	err := ctx.ShouldBindJSON(&tabungan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	result, err := c.tabunganService.CreateTabungan(ctx.Request.Context(), &tabungan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success create tabungan",
		"data": map[string]any{
			"tabungan": result,
		},
	})
}

func (c *TabunganControllerImpl) DetailTabungan(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	result, err := c.tabunganService.FindTabunganById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success get tabungan",
		"data": map[string]any{
			"tabungan": result,
		},
	})
}

func (c *TabunganControllerImpl) ListAllTabungan(ctx *gin.Context) {
	result, err := c.tabunganService.FindAllTabungan(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success get all tabungan",
		"data": map[string]any{
			"tabungan": result,
		},
	})
}

func (c *TabunganControllerImpl) UpdateTabungan(ctx *gin.Context) {
	if _, err := ctx.Cookie("jwt"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": err.Error(),
		})
		return
	}

	var tabungan models.Tabungan
	err := ctx.ShouldBindJSON(&tabungan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	if id, err := strconv.Atoi(ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	} else {
		tabungan.ID = uint(id)
	}

	result, err := c.tabunganService.UpdateTabungan(ctx.Request.Context(), &tabungan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success update tabungan",
		"data": map[string]any{
			"tabungan": result,
		},
	})
}

func (c *TabunganControllerImpl) DeleteTabungan(ctx *gin.Context) {
	if _, err := ctx.Cookie("jwt"); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	err = c.tabunganService.DeleteTabungan(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success delete tabungan",
	})
}
