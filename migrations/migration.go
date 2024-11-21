package main

import (
	"go-gin-udemy/infra"
	"go-gin-udemy/models"
)

func main() {
	infra.Initializer()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}