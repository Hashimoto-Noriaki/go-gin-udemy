package repositories

import (
	"errors"
	"go-gin-udemy/models"
)

// インターフェース IItemRepository の定義。
type IItemRepository interface {
	// repositoryが満たすべきメソッドを定義
	FindAll() (*[]models.Item, error)               // 商品情報のリスト（Item構造体のスライス）へのポインタ
	FindById(itemId int) (*models.Item, error) // itemIdの型をintに設定
	Create(newItem models.Item)(*models.Item,error)
}

// ItemMemoryRepositoryはメモリ上でデータを保持するための構造体
type ItemMemoryRepository struct {
	items []models.Item
}

// ItemMemoryRepositoryのインスタンスを生成して返すコンストラクタ関数。
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// FindAllメソッドは、ItemMemoryRepository構造体のメソッド。
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

// FindByIdメソッドは、指定されたIDのItemを返します
func (r *ItemMemoryRepository) FindById(itemId int) (*models.Item, error) {
	for _, v := range r.items {
		if int(v.ID) == itemId { // v.IDをintにキャストして比較
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}

// FindAllメソッドは、ItemMemoryRepository構造体のメソッド。
func (r *ItemMemoryRepository) Create(newItem models.Item)(*models.Item,error){
	newItem.ID = unit(len(r.items) + 1)
	r.items = append(r.items,newItem)
	return &newItem, nil
}
