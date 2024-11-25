package controllers

import (
    "go-gin-udemy/models"
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
    Delete(ctx *gin.Context)
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
    user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*models.User).ID
	
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	item, err := c.service.FindById(uint(itemId), userId)
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
    user, exists := ctx.Get("user")
    if !exists {
        ctx.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    userId := user.(*models.User).ID

    var input dto.CreateItemInput // dto パッケージを使用
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newItem, err := c.service.Create(input,userId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

// Update メソッドの実装
func (c *ItemController) Update(ctx *gin.Context) {
    func (c *ItemController) Update(ctx *gin.Context) {
        user, exists := ctx.Get("user")
        if !exists {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }
    
        userId := user.(*models.User).ID
    
        itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
            return
        }
    
        var input dto.UpdateItemInput
        if err := ctx.ShouldBindJSON(&input); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        updatedItem, err := c.service.Update(uint(itemId), userId, input)
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
}

func (c *ItemController) Delete(ctx *gin.Context) {
    itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
        return 
    }

    err = c.service.Delete(uint(itemId))    
    if err != nil {
        if err.Error() == "Item not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }
    ctx.Status(http.StatusOK)
}
