package services

import (
    "go-gin-udemy/models"
    "go-gin-udemy/repositories"
)

// サービスのインターフェース定義
type IItemServices interface {
    FindAll() (*[]models.Item, error)
    FindById(itemId int) (*models.Item, error) // itemIdの型をunitからintに修正
    Create(CreateItemInput dto.CreateItemInput) (*models.Item, error)
}

// サービスの構造体定義
type ItemServices struct {
    repository repositories.IItemRepository
}

// NewItemService関数の戻り値の型をIItemServicesに変更
func NewItemService(repository repositories.IItemRepository) IItemServices {
    return &ItemServices{repository: repository}  // 正しい構造体リテラルの書き方
}

// FindAllメソッドの定義
func (s *ItemServices) FindAll() (*[]models.Item, error) {
    return s.repository.FindAll()
}

// FindByIdメソッドの定義
func (s *ItemServices) FindById(itemId int) (*models.Item, error) { // itemIdの型をunitからintに修正
    return s.repository.FindById(itemId)
}

func (s *ItemServices) Create(CreateItemInput dto.CreateItemInput) (*models.Item, error) { // itemIdの型をunitからintに修正
    newItem := models.Item{
        Name: createItemInput.Name,
        Price: createItemInput.Price,
        Description: createItemInput.Description,
        SoldOut: false,
    }
    return s.repository.Create(newItem)
}