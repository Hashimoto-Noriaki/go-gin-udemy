package services

import (
    "go-gin-udemy/models"
    "go-gin-udemy/repositories"
)

type IItemServices interface {
    FindAll() (*[]models.Item, error)
    FindById(itemId unit) (*models.Item, error)
}

type ItemServices struct {
    repository repositories.IItemRepository
}

// NewItemService関数の戻り値の型をIItemServicesに変更
func NewItemService(repository repositories.IItemRepository) IItemServices {
    return &ItemServices{repository: repository}  // 正しい構造体リテラルの書き方
}

func (s *ItemServices) FindAll() (*[]models.Item, error) {
    return s.repository.FindAll()
}

func (s *ItemMemoryRepository) FindById(itemId unit)(*models.Item,error){
    return s.repository.FindById(itemId)
}
