package controllers

import "github.com/gin-gonic/gin"

type ResultController interface {
	HitungResult(ctx *gin.Context)
}
