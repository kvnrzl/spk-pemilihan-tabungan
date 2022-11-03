package controllers

import "github.com/gin-gonic/gin"

type AdminController interface {
	AdminRegister(r *gin.Context)
	AdminLogin(r *gin.Context)
	AdminLogout(r *gin.Context)
}
