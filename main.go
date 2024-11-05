package main

import (
    "go-gin-udemy/models"
    "go-gin-udemy/repositories"
    "go-gin-udemy/services"
    "go-gin-udemy/controllers"

    "github.com/gin-gonic/gin"
)

func main() {
    items := []models.Item{
        {ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
        {ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
        {ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
    }

    itemRepository := repositories.NewItemMemoryRepository(items)
    itemService := services.NewItemService(itemRepository)
    itemController := controllers.NewItemController(itemService)  // itemServiceを渡す

    r := gin.Default() // GinのデフォルトのHTTPルーターを作成
    r.GET("/items", itemController.FindAll) // エンドポイントを追加
    r.Run("localhost:8080")
}
