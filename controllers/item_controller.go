package controllers

import (
    "go-gin-udemy/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type IItemController interface {
    FindAll(ctx *gin.Context)
}

type ItemController struct {
    service services.IItemServices
}

func NewItemController(service services.IItemServices) IItemController {  // 名前を修正
    return &ItemController{service: service}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
    items, err := c.service.FindAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": items})
}
