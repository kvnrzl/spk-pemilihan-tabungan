package controllers

import "github.com/gin-gonic/gin"

type PresetKriteriaController interface {
	CreatePreset(ctx *gin.Context)
	FindFirstPreset(ctx *gin.Context)
	UpdatePreset(ctx *gin.Context)
}
