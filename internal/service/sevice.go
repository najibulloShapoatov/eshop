package service

import (
	"eshop/internal/domain"
	"eshop/internal/models"
	"eshop/internal/repository"
	"eshop/internal/service/auth"
	"eshop/internal/service/cart"
	"eshop/internal/service/cartitem"
	"eshop/internal/service/product"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user *models.CreateUser) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, string, error)
}

type Product interface {
	Create(product *models.Product) (int64, error)
	GetByID(id int64) (*domain.Product, error)
	GetList() ([]*domain.Product, error)
	Update(id int64, product *models.Product) error
	Delete(id int64) error
}
type Cart interface {
	GetList() ([]*domain.Cart, error)
	Get(userId int64) (*domain.Cart, error)
	SaveProductToCart(userId int64, input []models.CartItem) (*domain.Cart, error)
}

type CartItem interface {
	Save(*models.CartItem) (int64, error)
	GetList(cartID int64) ([]*domain.CartItem, error)
}

type Service struct {
	Authorization
	Product
	Cart
	CartItem
}

func NewService(repos *repository.Repository) *Service {
	productService := product.NewProductService(repos.Product)
	cartItemService := cartitem.NewCartItemService(repos.CartItem, repos.Product)
	cartService := cart.NewCartService(repos.Cart, repos.CartItem, cartItemService)
	return &Service{
		Authorization: auth.NewAuthService(repos.Auth),
		Product:       productService,
		Cart:          cartService,
		CartItem:      cartItemService,
	}
}
