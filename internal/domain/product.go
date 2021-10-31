package domain

import (
	"eshop/internal/models"
	"time"
)

type Product struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	Quantity    int       `json:"quantity" db:"quantity"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(input *models.Product) *Product {
	now := time.Now()
	return &Product{
		Name:        input.Name,
		Description: input.Description,
		IsActive:    input.IsActive,
		Quantity:    input.Quantity,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
