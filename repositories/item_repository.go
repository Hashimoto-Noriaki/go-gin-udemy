package repositories

import "go-gin-udemy/models";

type IItemRepository interface {
	//repositoryが満たすべきメソッドを定義
	FindAll() (*[]models.Item,error)
}

type ItemMemoryRepository struct {
	//メモリ上のデータソースとして
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemsMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item,error){
	return &r.items,nil
}
