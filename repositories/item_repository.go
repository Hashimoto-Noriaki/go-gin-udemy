package repositories

import "go-gin-udemy/models";
//インターフェース IItemRepository の定義。 IItemRepositoryインターフェースは、リポジトリが満たすべきメソッドを定義。この場合、商品データを全て取得するメソッドFindAllのみが定義
type IItemRepository interface {
	//repositoryが満たすべきメソッドを定義
	FindAll() (*[]models.Item,error)//*[]models.Item：商品情報のリスト（Item構造体のスライス）へのポインタ。
	FindById(itemId unit) (*models.Item, error)
}
//temMemoryRepositoryはメモリ上でデータを保持するための構造体。フィールドitemsに、商品情報を格納するスライス[]models.Itemが定義
type ItemMemoryRepository struct {
	items []models.Item
}
//ItemMemoryRepositoryのインスタンスを生成して返すコンストラクタ関数。
//itemsスライスを引数として受け取り、そのスライスをItemMemoryRepository構造体のフィールドに設定。
//この関数は、IItemRepositoryインターフェースを返すので、他のパッケージでIItemRepositoryインターフェースとして扱う
func NewItemMemoryRepository(items []models.Item) IItemRepository {
    return &ItemMemoryRepository{items: items}
}
//FindAllメソッドは、ItemMemoryRepository構造体のメソッド。
//構造体内のitemsスライスをポインタで返す。
//戻り値の型は*[]models.Item（Itemのリストへのポインタ）とerrorで、エラーは発生しないのでnilを返しています。
func (r *ItemMemoryRepository) FindAll() (*[]models.Item,error){
	return &r.items,nil
}

func (r *ItemMemoryRepository) FindAll()(*[]models.Item,error){
	return &r.items,nil
}
