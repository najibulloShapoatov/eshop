package cart

import (
	"eshop/internal/domain"
	"eshop/internal/models"
	"eshop/internal/repository"
	"eshop/internal/service/cartitem"

	"github.com/sirupsen/logrus"
)

type CartService struct {
	repo            repository.Cart
	cartItemRepo    repository.CartItem
	cartItemService *cartitem.CartItemService
}

func NewCartService(
	repo repository.Cart,
	cartItemRepo repository.CartItem,
	cartItemService *cartitem.CartItemService,
) *CartService {
	return &CartService{
		repo:            repo,
		cartItemService: cartItemService,
		cartItemRepo:    cartItemRepo,
	}
}

//Get one Cart or Create and Get
func (cartService *CartService) Get(userId int64) (*domain.Cart, error) {
	cart, err := cartService.repo.GetByUserID(userId)
	if err != nil {
		Cart := domain.NewCart(userId)
		cartService.repo.Create(Cart)
		cart, err = cartService.repo.GetByUserID(userId)
	}
	if err != nil {
		return nil, err
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	return cart, nil
}

//GetList  all Carts get method
func (cartService *CartService) GetList() ([]*domain.Cart, error) {
	carts, _ := cartService.repo.GetList()
	for _, item := range carts {
		item.Items, _ = cartService.cartItemService.GetList(item.ID)
	}
	return carts, nil
}

//SaveProductToCart  updating or adding product to cart
func (cartService *CartService) SaveProductToCart(userId int64, input []models.CartItem) (*domain.Cart, error) {
	cart, err := cartService.Get(userId)
	if err != nil {
		return nil, err
	}
	for _, item := range input {
		item.CartId = cart.ID
		_, err := cartService.cartItemService.Save(&item)
		logrus.Error(err)
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	return cart, nil
}
