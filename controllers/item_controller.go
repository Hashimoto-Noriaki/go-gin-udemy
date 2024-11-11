package controllers

import (
    "go-gin-udemy/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type IItemController interface {
    FindAll(ctx *gin.Context)
    FindById(ctx *gin.Context)
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

func (c *ItemController) FindById(ctx *gin.Context){
    itemId, err := strconv.ParseUnit(ctx.Params("id"),10,64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    item, err := c.service.FindById(unit(itemId))
    if err != nil {
        if err.Error() == "Item not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError,gin.H{"error": "unexpected error"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": item})
}