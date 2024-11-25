package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitempty,min=2"`
	Price       *uint   `json:"price" binding:"omitempty,min=1,max=999999"` // ポインタ型に変更
	Description *string `json:"description" binding:"omitempty"`           // ポインタ型に変更
	SoldOut     *bool   `json:"soldOut"`                                    // そのままポインタ型
}
