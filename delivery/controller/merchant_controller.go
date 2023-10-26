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

type MerchantController struct {
	Router        *gin.Engine
	merchantUsecase usecase.MerchantUsecase
}

func (c *MerchantController) Create(ctx *gin.Context) {
	var merchant model.Merchant
	merchant.Id = common.GenerateUUID()
	err := ctx.ShouldBindJSON(&merchant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.merchantUsecase.Create(merchant)
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

func (c *MerchantController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	name := ctx.Query("name")
	paginationParam := dto.PaginationQueryParam{
		Page:  page,
		Limit: limit,
	}
	// merchants, err := c.merchantUsecase.List()
	merchants, pagination, err := c.merchantUsecase.Pagination(paginationParam, name)
	fmt.Println(merchants)
	if len(merchants) == 0 {
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
		"data":      merchants,
		"data_page": pagination,
	})
}

func (c *MerchantController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := c.merchantUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   merchant,
	})
}

func (c *MerchantController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var merchant model.Merchant
	merchant.Id = id

	_, err := c.merchantUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&merchant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.merchantUsecase.Update(merchant)
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

func (c *MerchantController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.merchantUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.merchantUsecase.Delete(id)
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

func NewMerchantController(router *gin.Engine, merchantUsecase usecase.MerchantUsecase) *MerchantController {
	controller := &MerchantController{
		Router:        router,
		merchantUsecase: merchantUsecase,
	}

	routerGroup := controller.Router.Group("/api/v1")
	routerGroup.GET("/merchant", middleware.AuthMiddleware(), controller.List)
	routerGroup.GET("/merchant/:id", middleware.AuthMiddleware(), controller.Get)
	routerGroup.POST("/merchant", middleware.AuthMiddleware(), controller.Create)
	routerGroup.PUT("/merchant/:id", middleware.AuthMiddleware(), controller.Update)
	routerGroup.DELETE("/merchant/:id", middleware.AuthMiddleware(),controller.Delete)
	return controller
}
