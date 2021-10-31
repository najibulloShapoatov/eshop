package repository

import (
	"eshop/internal/domain"
	"eshop/internal/repository/sql"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth
	Product
	Cart
	CartItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     sql.NewAuth(db),
		Product:  sql.NewProduct(db),
		Cart:     sql.NewCart(db),
		CartItem: sql.NewCartItem(db),
	}
}

type Auth interface {
	CreateUser(user *domain.User) (int64, error)
	GetUser(username, password string) (*domain.User, error)
}
type Product interface {
	Create(product *domain.Product) (int64, error)
	GetByID(id int64) (*domain.Product, error)
	GetList() ([]*domain.Product, error)
	Update(id int64, form *domain.Product) error
	Delete(id int64) error
}
type Cart interface {
	Create(cart *domain.Cart) (int64, error)
	GetByUserID(id int64) (*domain.Cart, error)
	GetList() ([]*domain.Cart, error)
	Delete(id int64) error
}

type CartItem interface {
	Create(cartItem *domain.CartItem) (int64, error)
	GetByID(id int64) (*domain.CartItem, error)
	GetByProductID(id int64, productId int64) (*domain.CartItem, error)
	GetList(cartID int64) ([]*domain.CartItem, error)
	Update(id int64, form *domain.CartItem) error
}
