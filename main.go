package main

import (
    "go-gin-udemy/repositories"
    "go-gin-udemy/services"
    "go-gin-udemy/controllers"
    "go-gin-udemy/infra"
    "go-gin-udemy/middlewares"

    "github.com/gin-gonic/gin"
)

func main() {
    infra.Initializer()
    db := infra.SetupDB()

    // Repositories
    itemRepository := repositories.NewItemRepository(db)
    authRepository := repositories.NewAuthRepository(db)

    // Services
    itemService := services.NewItemService(itemRepository)
    authService := services.NewAuthService(authRepository)

    // Controllers
    itemController := controllers.NewItemController(itemService)
    authController := controllers.NewAuthController(authService)

    // Gin Engine
    r := gin.Default()

    // /items グループ化
    itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService)) // AuthMiddleware を正しく適用
    {
        itemRouterWithAuth.POST("", itemController.Create) // 認証が必要なルート
        itemRouterWithAuth.PUT("/:id", itemController.Update)
        itemRouterWithAuth.DELETE("/:id", itemController.Delete)
    }

    itemRouter := r.Group("/items")
    {
        itemRouter.GET("", itemController.FindAll) // 認証不要なルート
        itemRouter.GET("/:id", itemController.FindById)
    }

    // /auth グループ化
    authRouter := r.Group("/auth")
    {
        authRouter.POST("/signup", authController.Signup)
        authRouter.POST("/login", authController.Login)
    }

    // サーバ起動
    r.Run("localhost:8080")
}
