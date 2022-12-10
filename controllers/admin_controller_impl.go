package controllers

import (
	"net/http"
	model "project_spk_pemilihan_tabungan/models"
	"project_spk_pemilihan_tabungan/services"

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

	result, err := c.adminService.AdminCreate(r.Request.Context(), admin.Username, admin.Password)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "register success",
		"data": map[string]any{
			"admin": result,
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
		r.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	// cookie := http.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(config.JWT_EXPIRE_DURATION),
	// 	Path:     "/", // cookie will be available on all pages
	// 	HttpOnly: true,
	// 	// SameSite: http.SameSiteNoneMode,
	// 	// Secure:   true,
	// }

	// r.SetCookie(cookie.Name, cookie.Value, 86400, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	// r.Writer.Header().Set("Set-Cookie", cookie.String())
	// http.SetCookie(r.Writer, &cookie)

	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data": map[string]any{
			"username": admin.Username,
			"token":    token,
		},
	})
}

func (c *AdminControllerImpl) AdminLogout(r *gin.Context) {
	// auth := r.Request.Header.Get("Authorization")

	// _, err := r.Cookie("jwt")
	// if err != nil {
	// 	r.JSON(http.StatusBadRequest, gin.H{
	// 		"code":  http.StatusBadRequest,
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// cookie := http.Cookie{
	// 	Name:    "jwt",
	// 	Value:   "",
	// 	Expires: time.Now().Add(-time.Hour),
	// 	Path:    "/", // cookie will be available on all pages
	// 	// HttpOnly: true,
	// 	// 	// SameSite: http.SameSiteNoneMode,
	// 	// 	// Secure:   true,
	// }
	// r.Writer.Header().Set("Set-Cookie", cookie.String())
	// http.SetCookie(r.Writer, &cookie)

	// r.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	// if auth != "" {
	// }
	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Logout Success",
	})
}
