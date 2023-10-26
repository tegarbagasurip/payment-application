package controller

import (
	"payment-application/delivery/middleware"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/usecase"
	"payment-application/utils/common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileController struct { // Updated struct name
	Router        *gin.Engine
	profileUsecase usecase.ProfileUsecase // Updated type name
}

func (c *ProfileController) Create(ctx *gin.Context) {
	var profile model.Profile // Updated variable name
	profile.Id = common.GenerateUUID()
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.profileUsecase.Create(profile) // Updated function name
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (c *ProfileController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	name := ctx.Query("name")
	paginationParam := dto.PaginationQueryParam{
		Page:  page,
		Limit: limit,
	}
	// profiles, err := c.profileUsecase.List()
	profiles, pagination, err := c.profileUsecase.Pagination(paginationParam, name)
	fmt.Println(profiles)
	if len(profiles) == 0 {
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
		"status":    "success",
		"data":      profiles,
		"data_page": pagination,
	})
}

func (c *ProfileController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	profile, err := c.profileUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   profile,
	})
}

func (c *ProfileController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var profile model.Profile // Updated variable name
	profile.Id = id

	_, err := c.profileUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.profileUsecase.Update(profile) // Updated function name
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


func NewProfileController(router *gin.Engine, profileUsecase usecase.ProfileUsecase) *ProfileController {
	controller := &ProfileController{ // Updated struct name
		Router:        router,
		profileUsecase: profileUsecase, // Updated variable name
	}

	routerGroup := controller.Router.Group("/api/v1")
	routerGroup.GET("/profile", middleware.AuthMiddleware(), controller.List) 
	routerGroup.GET("/profile/:id", middleware.AuthMiddleware(), controller.Get) 
	routerGroup.POST("/profile", middleware.AuthMiddleware(), controller.Create) 
	routerGroup.PUT("/profile/:id", middleware.AuthMiddleware(), controller.Update) 
	// routerGroup.DELETE("/profile/:id", controller.Delete) // Updated route
	return controller
}
