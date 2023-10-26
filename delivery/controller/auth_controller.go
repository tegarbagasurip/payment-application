package controller

import (
	"net/http"
	"payment-application/delivery/middleware"
	"payment-application/model"
	"payment-application/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router      *gin.Engine
	authUsecase usecase.AuthUsecase
}

func (c *AuthController) createHandler(ctx *gin.Context) {
	var userCredential model.UserCredential
	err := ctx.ShouldBindJSON(&userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	token, err := c.authUsecase.Login(userCredential.Username, userCredential.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

func (c *AuthController) logoutHandler(ctx *gin.Context) {
	user, exists := ctx.Get("user") // Mendapatkan informasi pengguna yang login dari middleware

	if exists {
	

		// Contoh: Hapus token dari basis data
		userID := user.(model.UserCredential).Id
		err := c.authUsecase.Logout(userID)

		// Jika ada kesalahan saat logout, Anda dapat merespons dengan pesan kesalahan.
		if err != nil {
		    ctx.JSON(http.StatusInternalServerError, gin.H{
		        "status":  "fail",
		        "message": err.Error(),
		    })
		    return
		}

		// Jika logout berhasil, Anda dapat memberi respon dengan pesan berhasil logout.
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Logout berhasil",
		})
	} else {
		// Jika tidak ada pengguna yang login (token tidak ditemukan), Anda dapat memberi respon dengan pesan sesi sudah logout.
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Anda sudah logout atau sesi tidak ditemukan.",
		})
	}
}



func NewAuthController(r *gin.Engine, authUsecase usecase.AuthUsecase) {
	controller := AuthController{
		router:      r,
		authUsecase: authUsecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/login", controller.createHandler)
	rg.POST("/logout", middleware.AuthMiddleware(), controller.logoutHandler)
}
