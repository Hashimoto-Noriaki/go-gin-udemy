package main

import "github.com/gin-gonic/gin"

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1",Price: 1000, Description: "説明1",SoldOut: false},
		{ID: 2, Name: "商品2",Price: 2000, Description: "説明2",SoldOut: true},
		{ID: 3, Name: "商品3",Price: 3000, Description: "説明3",SoldOut: false}
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	ItemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController(ItemServer)

	r := gin.Default()//これはHTTPリクエストを処理
	r.GET("/items", ItemController.FindAll)//ルーターにクライアントのリクエスト先となるエンドポイントを追加
	r.Run("localhost:8080")
}
