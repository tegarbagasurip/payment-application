package controller

import (
	"payment-application/model"
	"payment-application/usecase"
	"net/http"

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

func NewAuthController(r *gin.Engine, authUsecase usecase.AuthUsecase) {
	controller := AuthController{
		router:      r,
		authUsecase: authUsecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/login", controller.createHandler)
}
