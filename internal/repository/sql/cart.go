package sql

import (
	"eshop/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Cart struct {
	db *sqlx.DB
}

func NewCart(db *sqlx.DB) *Cart {
	return &Cart{db: db}
}

func (c *Cart) Create(Cart *domain.Cart) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING id", cartsTable)
	row := c.db.QueryRow(query, Cart.UserID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Cart) GetByUserID(id int64) (*domain.Cart, error) {
	query := fmt.Sprintf("select * from %s where user_id =$1", cartsTable)
	var Cart domain.Cart
	err := c.db.Get(&Cart, query, id)
	return &Cart, err
}

func (c *Cart) GetList() ([]*domain.Cart, error) {
	query := fmt.Sprintf("select * from %s ", cartsTable)
	var Carts []*domain.Cart
	err := c.db.Select(&Carts, query)
	return Carts, err
}

func (c *Cart) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id =$1", cartsTable)
	_, err := c.db.Exec(query, id)
	return err
}
