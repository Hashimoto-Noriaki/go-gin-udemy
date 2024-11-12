package services

import (
    "go-gin-udemy/models"
    "go-gin-udemy/repositories"
    "go-gin-udemy/dto" // dtoパッケージのインポートを追加
)

// サービスのインターフェース定義
type IItemServices interface {
    FindAll() (*[]models.Item, error)
    FindById(itemId int) (*models.Item, error) 
    Create(createItemInput dto.CreateItemInput) (*models.Item, error) // 修正: 引数名を小文字開始に変更
}

// サービスの構造体定義
type ItemServices struct {
    repository repositories.IItemRepository
}

// NewItemService関数
func NewItemService(repository repositories.IItemRepository) IItemServices {
    return &ItemServices{repository: repository}
}

// FindAllメソッドの定義
func (s *ItemServices) FindAll() (*[]models.Item, error) {
    return s.repository.FindAll()
}

// FindByIdメソッドの定義
func (s *ItemServices) FindById(itemId int) (*models.Item, error) {
    return s.repository.FindById(itemId)
}

// Createメソッドの定義
func (s *ItemServices) Create(createItemInput dto.CreateItemInput) (*models.Item, error) { // 引数名を修正
    newItem := models.Item{
        Name:        createItemInput.Name,
        Price:       createItemInput.Price,
        Description: createItemInput.Description,
        SoldOut:     false,
    }
    return s.repository.Create(newItem)
}
