package product

import (
	"eshop/internal/domain"
	"eshop/internal/models"
	"eshop/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (productService *ProductService) Create(input *models.Product) (int64, error) {
	product := domain.NewProduct(input)
	return productService.repo.Create(product)
}
func (productService *ProductService) GetByID(id int64) (*domain.Product, error) {
	return productService.repo.GetByID(id)
}
func (productService *ProductService) GetList() ([]*domain.Product, error) {
	return productService.repo.GetList()
}
func (productService *ProductService) Update(id int64, input *models.Product) error {
	product := domain.NewProduct(input)
	return productService.repo.Update(id, product)
}
func (productService *ProductService) Delete(id int64) error {
	return productService.repo.Delete(id)
}
