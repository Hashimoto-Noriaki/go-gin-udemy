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

    itemRouterWithAuth := r.Group("/items",middlewares.AuthMiddleware(authService))

    // /items のルートをグループ化
    itemRouter := r.Group("/items")
    {
        itemRouter.GET("", itemController.FindAll)
        itemRouter.GET("/:id", itemController.FindById)
        itemRouterWithAuth.POST("", itemController.Create)
        itemRouter.PUT("/:id", itemController.Update)
        itemRouter.DELETE("/:id", itemController.Delete)
    }

    // /auth のルートをグループ化
    authRouter := r.Group("/auth")
    {
        authRouter.POST("/signup", authController.Signup)
        authRouter.POST("/login", authController.Login)
    }

    r.Run("localhost:8080")
}
