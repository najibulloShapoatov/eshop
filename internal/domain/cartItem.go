package domain

type CartItem struct {
	ID        int64    `json:"id" db:"id"`
	CartID    int64    `json:"cart_id" db:"cart_id"`
	Quantity  int64    `json:"quantity" db:"quantity"`
	ProductID int64    `json:"product_id" db:"product_id"`
	Product   *Product `json:"product" db:"-"`
}

func NewCartItem(CartId int64, ProductId int64, qty int64) *CartItem {
	return &CartItem{
		CartID:    CartId,
		ProductID: ProductId,
		Quantity:  qty,
	}
}
