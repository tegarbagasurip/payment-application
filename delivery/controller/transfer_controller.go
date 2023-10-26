package controller

import (
	"fmt"
	"net/http"
	"payment-application/delivery/middleware"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	Router          *gin.Engine
	transferUsecase usecase.TransferUsecase
}

func (c *TransferController) Create(ctx *gin.Context) {
	var transfer model.Transfer
	err := ctx.ShouldBindJSON(&transfer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	// Panggil use case untuk melakukan transfer
	err = c.transferUsecase.TransferBalance(transfer)
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

func (c *TransferController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	description := ctx.Query("description")
	paginationParam := dto.PaginationQueryParam{
		Page:  page,
		Limit: limit,
	}
	transfers, pagination, err := c.transferUsecase.Pagination(paginationParam, description)
	fmt.Println(transfers)
	if len(transfers) == 0 {
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
		"data":      transfers,
		"data_page": pagination,
	})
}

func NewTransferController(router *gin.Engine, transferUsecase usecase.TransferUsecase) *TransferController {
	controller := &TransferController{
		Router:          router,
		transferUsecase: transferUsecase,
	}

	routerGroup := controller.Router.Group("/api/v1")
	routerGroup.GET("/transfer", middleware.AuthMiddleware(), controller.List)
	routerGroup.POST("/transfer", middleware.AuthMiddleware(), controller.Create)
	return controller
}
