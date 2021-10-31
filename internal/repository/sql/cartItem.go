package sql

import (
	"eshop/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CartItem struct {
	db *sqlx.DB
}

func NewCartItem(db *sqlx.DB) *CartItem {
	return &CartItem{db: db}
}

func (c *CartItem) Create(CartItem *domain.CartItem) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (cart_id, product_id, quantity) values ($1,$2,$3) RETURNING id", cartItemsTable)
	row := c.db.QueryRow(query, CartItem.CartID, CartItem.ProductID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CartItem) GetByID(id int64) (*domain.CartItem, error) {
	query := fmt.Sprintf("select * from %s where id =$1", cartItemsTable)
	var CartItem domain.CartItem
	err := c.db.Get(&CartItem, query, id)
	return &CartItem, err
}

func (c *CartItem) GetList(id int64) ([]*domain.CartItem, error) {
	query := fmt.Sprintf("select * from %s where cart_id=$1", cartItemsTable)
	var CartItems []*domain.CartItem
	err := c.db.Select(&CartItems, query, id)
	return CartItems, err
}

func (c *CartItem) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id =$1", cartItemsTable)
	_, err := c.db.Exec(query, id)
	return err
}

func (c *CartItem) Update(id int64, form *domain.CartItem) error {
	query := fmt.Sprintf("update %s set quantity=$1 where id=$2", cartItemsTable)
	_, err := c.db.Exec(query, form.Quantity, id)
	return err
}
