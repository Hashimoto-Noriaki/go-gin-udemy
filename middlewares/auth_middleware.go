package middlewares

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "go-gin-udemy/services"
)

func AuthMiddleware(authService services.IAuthServices) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        header := ctx.GetHeader("Authorization")
        if header == "" {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        if !strings.HasPrefix(header, "Bearer") {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(header, "Bearer")
        user, err := authService.GetUserFromToken(tokenString)
        if err != nil {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        ctx.Set("user", user)
        ctx.Next()
    }
}
