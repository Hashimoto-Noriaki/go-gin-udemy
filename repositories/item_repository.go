package repositories

import (
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

// NewItemRepository 関数（関数名が間違っている可能性があるので確認）
func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

// それぞれのメソッドを実装
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (r *ItemRepository) Delete(itemId uint) error {
	result := r.db.Delete(&models.Item{}, itemId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
