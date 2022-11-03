package controllers

import "github.com/gin-gonic/gin"

type TabunganController interface {
	CreateTabungan(ctx *gin.Context)
	DetailTabungan(ctx *gin.Context)
	ListAllTabungan(ctx *gin.Context)
	UpdateTabungan(ctx *gin.Context)
	DeleteTabungan(ctx *gin.Context)
}
