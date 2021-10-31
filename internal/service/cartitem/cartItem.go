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

//Save  add or update product from cart
func (cartItemService *CartItemService) Save(input *models.CartItem) (int64, error) {
	product, err := cartItemService.repoProduct.GetByID(input.ProductId)
	if err != nil {
		return 0, err
	}
	if product.Quantity < int(input.Quantity) {
		return 0, errors.New("quantities greater than stock quantities")
	}
	cartItem, err := cartItemService.repo.GetByProductID(input.CartId, input.ProductId)
	if err != nil {
		CartItem := domain.NewCartItem(input.CartId, input.ProductId, input.Quantity)
		return cartItemService.repo.Create(CartItem)
	}
	cartItem.Quantity = input.Quantity

	err = cartItemService.repo.Update(cartItem.ID, cartItem)
	if err != nil {
		return 0, err
	}
	return cartItem.ID, nil
}

//Get All itemsfor cart
func (cartItemService *CartItemService) GetList(id int64) ([]*domain.CartItem, error) {
	items, err := cartItemService.repo.GetList(id)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		item.Product, _ = cartItemService.repoProduct.GetByID(item.ProductID)
	}

	return items, nil
}
