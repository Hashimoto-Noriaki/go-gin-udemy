package main

import (
    // "go-gin-udemy/models"
    "go-gin-udemy/repositories"
    "go-gin-udemy/services"
    "go-gin-udemy/controllers"
    "go-gin-udemy/infra"

    "github.com/gin-gonic/gin"
)

func main() {
    infra.Initializer()
    db := infra.SetupDB()

    // items := []models.Item{
    //     {ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
    //     {ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
    //     {ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
    // }

    // itemRepository := repositories.NewItemMemoryRepository(items)
    itemRepository := repositories.NewItemRepository(db)  // ここが正しい関数名であることを確認
    itemService := services.NewItemService(itemRepository)
    itemController := controllers.NewItemController(itemService)

    authRepository := repositories.NewAuthRepository(db)
    authService := services.NewAuthService(authRepository)
    authController := controllers.NewAuthController(authService)

    r := gin.Default() // GinのデフォルトのHTTPルーターを作成
    itemRouter := r.Group("/items")
    authRouter := r.Group("/auth")

    r.GET("", itemController.FindAll) // エンドポイントを追加
    r.GET("/:id",itemController.FindById)
    r.POST("",itemController.Create)
    r.PUT("/:id",itemController.Update)
    r.DELETE("/:id",itemController.Delete)

    authRouter.POST("/signup", authController.Signup)
    r.Run("localhost:8080")
}
