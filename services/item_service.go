package services

import (
	"go-gin-udemy/models"
	"go-gin-udemy/repositories"
)

type IItemServices interface {
	FindAll() (*[]models.Item,error)
} 

type ItemServices struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemRepository {
	return &ItemService(repository: repository)
}

func  (s *ItemService) FindAll() (*[]models.Item,error) {
	return s.repositories.FindAll()
}
