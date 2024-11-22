package main

import (
    "go-gin-udemy/repositories"
    "go-gin-udemy/services"
    "go-gin-udemy/controllers"
    "go-gin-udemy/infra"

    "github.com/gin-gonic/gin"
)

func main() {
    infra.Initializer()
    db := infra.SetupDB()

    itemRepository := repositories.NewItemRepository(db)
    itemService := services.NewItemService(itemRepository)
    itemController := controllers.NewItemController(itemService)

    authRepository := repositories.NewAuthRepository(db)
    authService := services.NewAuthService(authRepository)
    authController := controllers.NewAuthController(authService)

    r := gin.Default()

    // /items のルートをグループ化
    itemRouter := r.Group("/items")
    {
        itemRouter.GET("", itemController.FindAll)     // itemRouter を利用
        itemRouter.GET("/:id", itemController.FindById)
        itemRouter.POST("", itemController.Create)
        itemRouter.PUT("/:id", itemController.Update)
        itemRouter.DELETE("/:id", itemController.Delete)
    }

    // /auth のルートをグループ化
    authRouter := r.Group("/auth")
    {
        authRouter.POST("/signup", authController.Signup)
    }

    r.Run("localhost:8080")
}
