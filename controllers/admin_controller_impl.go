package controllers

import (
	"net/http"
	"project_spk_pemilihan_tabungan/config"
	model "project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminControllerImpl struct {
	adminService services.AdminService
}

func NewAdminController(service services.AdminService) AdminController {
	return &AdminControllerImpl{
		adminService: service,
	}
}

func (c *AdminControllerImpl) AdminRegister(r *gin.Context) {
	var admin model.Admin

	if err := r.ShouldBindJSON(&admin); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	newAdmin, err := c.adminService.AdminCreate(r.Request.Context(), admin.Username, admin.Password)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]any{
			"admin": newAdmin,
		},
	})
}

func (c *AdminControllerImpl) AdminLogin(r *gin.Context) {
	var admin model.Admin

	if err := r.ShouldBindJSON(&admin); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	token, err := c.adminService.AdminLogin(r.Request.Context(), admin.Username, admin.Password)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(config.JWT_EXPIRE_DURATION),
	}

	// r.SetCookie(cookie.Name, cookie.Value, 86400, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	http.SetCookie(r.Writer, &cookie)

	r.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]any{
			"token":   token,
			"message": "Login Success",
		},
	})
}

func (c *AdminControllerImpl) AdminLogout(r *gin.Context) {
	_, err := r.Cookie("jwt")
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	// r.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	http.SetCookie(r.Writer, &cookie)

	r.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]any{
			"message": "Logout Success",
		},
	})
}
