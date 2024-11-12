package controllers

import (
    "go-gin-udemy/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// コントローラのインターフェース定義
type IItemController interface {
    FindAll(ctx *gin.Context)
    FindById(ctx *gin.Context)
    Create(ctx *gin.Context)
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
    // URL パラメータ "id" を取得
    itemId, err := strconv.Atoi(ctx.Param("id")) // 修正: ctx.Param を使用し strconv.Atoi に変更
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    // サービスからアイテムを取得
    item, err := c.service.FindById(itemId) // 修正: itemIdの型をintに
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

func (c *ItemController) Create(ctx *gin.Context) {
    var input dto.CreateItemInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
        return
    }
    newItem, err := c.service.Create(input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}
