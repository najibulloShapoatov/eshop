package models

type CartItem struct {
	CartId    int64 `json:"-"`
	ProductId int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"qty" binding:"required"`
}
