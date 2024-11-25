package middlewares

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "go-gin-udemy/services"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc { // IAuthService に修正
    return func(ctx *gin.Context) {
        header := ctx.GetHeader("Authorization")
        if header == "" {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        if !strings.HasPrefix(header, "Bearer ") { // "Bearer " スペースを考慮
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer"))
        user, err := authService.GetUserFromToken(tokenString)
        if err != nil {
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        ctx.Set("user", user) // コンテキストにユーザーを設定
        ctx.Next()
    }
}
