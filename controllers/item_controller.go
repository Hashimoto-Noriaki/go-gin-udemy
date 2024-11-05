package controllers

import (
	"go-gin-udemy/services"
	"net-http"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	FindAll(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemMemoryController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (C *ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServiceError,gin.H{"error":"Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"data": items})
} 