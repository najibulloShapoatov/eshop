package cartitem

import (
	"errors"
	"eshop/internal/domain"
	"eshop/internal/models"
	"eshop/internal/repository"
)

type CartItemService struct {
	repo        repository.CartItem
	repoProduct repository.Product
}

func NewCartItemService(repo repository.CartItem, repoProduct repository.Product) *CartItemService {
	return &CartItemService{
		repo:        repo,
		repoProduct: repoProduct,
	}
}

func (cartItemService *CartItemService) Create(input *models.CartItem) (int64, error) {
	product, err := cartItemService.repoProduct.GetByID(input.ProductId)
	if err != nil {
		return 0, err
	}
	if product.Quantity < int(input.Quantity) {
		return 0, errors.New("quantities greater than stock quantities")
	}
	CartItem := domain.NewCartItem(input.CartId, input.ProductId, input.Quantity)
	return cartItemService.repo.Create(CartItem)
}

func (cartItemService *CartItemService) GetByID(id int64) (*domain.CartItem, error) {
	return cartItemService.repo.GetByID(id)
}

func (cartItemService *CartItemService) GetList(id int64) ([]*domain.CartItem, error) {
	return cartItemService.repo.GetList(id)
}



func (cartItemService *CartItemService) Delete(id int64) error {
	return cartItemService.repo.Delete(id)
}
