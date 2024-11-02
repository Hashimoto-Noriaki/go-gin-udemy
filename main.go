package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()//これはHTTPリクエストを処理
	r.GET("/ping", func(c *gin.Context) {//ルーターにクライアントのリクエスト先となるエンドポイントを追加
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080")
}
