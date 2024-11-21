package repositories

import (
	"errors"
	"go-gin-udemy/models"
	"gorm.io/gorm"
)

// IItemRepository インターフェースの定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemRepository struct {
	db *gorm.DB
}

// NewItemRepository 関数
func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

// DBから全データ取得
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// IDに基づくデータ取得
func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { // gorm.ErrRecordNotFoundをチェック
		return nil, errors.New("Item not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// データの作成
func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// データの更新の実装
func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

// データの削除
func (r *ItemRepository) Delete(itemId uint) error {
	result := r.db.Delete(&models.Item{}, itemId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { // 存在しない場合のエラーチェック
		return errors.New("Item not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
