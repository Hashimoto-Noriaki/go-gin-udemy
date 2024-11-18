package controllers

import (
    "go-gin-udemy/services"
    "go-gin-udemy/dto"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// コントローラのインターフェース定義
type IItemController interface {
    FindAll(ctx *gin.Context)
    FindById(ctx *gin.Context)
    Create(ctx *gin.Context)
    Update(ctx *gin.Context)
}

// コントローラの構造体定義
type ItemController struct {
    service services.IItemServices
}

// 新しいコントローラを作成
func NewItemController(service services.IItemServices) IItemController {
    return &ItemController{service: service}
}

// FindAll メソッドの実装
func (c *ItemController) FindAll(ctx *gin.Context) {
    items, err := c.service.FindAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// FindById メソッドの実装
func (c *ItemController) FindById(ctx *gin.Context) {
    // itemIdをint型で取得して、uintにキャスト
    itemId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    // uintにキャスト
    itemIdUint := uint(itemId)

    item, err := c.service.FindById(itemIdUint)
    if err != nil {
        if err.Error() == "Item not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": item})
}

// Create メソッドの実装
func (c *ItemController) Create(ctx *gin.Context) {
    var input dto.CreateItemInput // dto パッケージを使用
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newItem, err := c.service.Create(input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

// Update メソッドの実装
func (c *ItemController) Update(ctx *gin.Context) {
    // itemIdをuint型で取得
    itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    var input dto.UpdateItemInput // dto パッケージを使用
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // uintにキャスト
    updatedItem, err := c.service.Update(uint(itemId), input)
    if err != nil {
        if err.Error() == "Item not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}
