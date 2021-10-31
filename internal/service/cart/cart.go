package cart

import (
	"eshop/internal/domain"
	"eshop/internal/models"
	"eshop/internal/repository"
	"eshop/internal/service/cartitem"
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

func (cartService *CartService) Create(UserId int64) (int64, error) {
	cart, err := cartService.repo.GetByUserID(UserId)
	if err != nil {
		Cart := domain.NewCart(UserId)
		return cartService.repo.Create(Cart)
	}

	return cart.ID, nil
}

func (cartService *CartService) GetByID(id int64) (*domain.Cart, error) {
	cart, err := cartService.repo.GetByUserID(id)
	if err != nil {
		return nil, err
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	return cart, nil
}

func (cartService *CartService) GetList() ([]*domain.Cart, error) {
	carts, _ := cartService.repo.GetList()
	for _, item := range carts {
		item.Items, _ = cartService.cartItemRepo.GetList(item.ID)

	}
	return carts, nil
}

func (cartService *CartService) AddProductToCart(userId int64, input []models.CartItem) (*domain.Cart, error) {
	cart, err := cartService.GetByID(userId)
	if err != nil {
		return nil, err
	}
	for _, item := range input {
		item.CartId = cart.ID
		cartService.cartItemService.Create(&item)
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	return cart, nil
}

func (cartService *CartService) DeleteProductFromCart(userId int64, productId int64, qty int64) (*domain.Cart, error) {
	cart, err := cartService.GetByID(userId)
	if err != nil {
		return nil, err
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	for _, item := range cart.Items {
		if item.ProductID == productId {
			if item.Quantity <= qty {
				cartService.cartItemService.Delete(item.ID)
			} else {
				item.Quantity = item.Quantity - qty
				cartService.cartItemRepo.Update(item.ID, item)
			}
		}
	}
	return cartService.GetByID(userId)
}

func (cartService *CartService) Delete(userId int64) error {
	cart, err := cartService.repo.GetByUserID(userId)
	if err != nil {
		return err
	}
	cart.Items, _ = cartService.cartItemService.GetList(cart.ID)
	for _, item := range cart.Items {
		cartService.cartItemService.Delete(item.ID)
	}
	return cartService.repo.Delete(cart.ID)
}
