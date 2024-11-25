package middlewares

import "go-gin-udemy/services"

func AuthMiddleware(authService services.IAuthServices) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        
    }
}