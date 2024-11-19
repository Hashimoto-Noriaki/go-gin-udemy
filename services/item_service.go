package services

import (
    "go-gin-udemy/models"
    "go-gin-udemy/repositories"
    "go-gin-udemy/dto" // dtoパッケージのインポートを追加
)

// サービスのインターフェース定義
type IItemServices interface {
    FindAll() (*[]models.Item, error)
    FindById(itemId uint) (*models.Item, error) // 引数を uint に修正
    Create(createItemInput dto.CreateItemInput) (*models.Item, error)
    Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
    Delete(itemId uint )error
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
func (s *ItemServices) FindById(itemId uint) (*models.Item, error) {
    return s.repository.FindById(itemId)
}

// Createメソッドの定義
func (s *ItemServices) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
    newItem := models.Item{
        Name:        createItemInput.Name,
        Price:       createItemInput.Price,
        Description: createItemInput.Description,
        SoldOut:     false,
    }
    return s.repository.Create(newItem)
}

// Updateメソッドの定義
func (s *ItemServices) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
    // itemId を使ってアイテムを取得
    targetItem, err := s.FindById(itemId)
    if err != nil {
        return nil, err
    }

    // Name の更新
    if updateItemInput.Name != nil {
        targetItem.Name = *updateItemInput.Name
    }

    // Price の更新
    if updateItemInput.Price != 0 { // nil チェックを適切に行う
        targetItem.Price = updateItemInput.Price
    }

    // Description の更新
    if updateItemInput.Description != "" { // nil チェックを適切に行う
        targetItem.Description = updateItemInput.Description
    }

    // SoldOut の更新
    if updateItemInput.SoldOut != nil { // ポインタ型で値が nil でないことを確認
        targetItem.SoldOut = *updateItemInput.SoldOut
    }

    // 更新されたアイテムを返す
    return s.repository.Update(*targetItem)
}

func (s *ItemServices) Delete(itemId uint) error {
    return s.repository.Delete(itemId)
}
