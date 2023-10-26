package controller

import (
	"payment-application/delivery/middleware"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/usecase"
	"payment-application/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router      *gin.Engine
	userUsecase usecase.UserUsecase
}

func (c *UserController) createHandler(ctx *gin.Context) {
	var userCredential model.UserCredential
	userCredential.Id = common.GenerateUUID()
	err := ctx.ShouldBindJSON(&userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.userUsecase.Register(userCredential)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (c *UserController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	users, paging, err := c.userUsecase.FindAllUser(dto.PaginationQueryParam{Page: page, Limit: limit})

	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "data not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
		"paging": paging,
	})
}

func (c *UserController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userUsecase.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (c *UserController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var userCredential model.UserCredential
	userCredential.Id = id

	_, err := c.userUsecase.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.userUsecase.UpdateUser(userCredential)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (c *UserController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.userUsecase.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.userUsecase.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (c *UserController) updatePasswordHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var userCredential model.UserCredential
	userCredential.Id = id

	_, err := c.userUsecase.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&userCredential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.userUsecase.UpdatePassword(userCredential)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func NewUserController(r *gin.Engine, userUsecase usecase.UserUsecase) {
	controller := UserController{
		router:      r,
		userUsecase: userUsecase,
	}
	rg := r.Group("/api/v1", middleware.AuthMiddleware())
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	rg.GET("/users/:id", controller.getHandler)
	rg.PUT("/users/:id", controller.updateHandler)
	rg.DELETE("/users/:id", controller.deleteHandler)
	rg.PUT("/users/:id/password", controller.updatePasswordHandler)
}
