package models

type Product struct {
	Name        string `json:"name" binding:"required" `
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"is_active"`
	Quantity    int    `json:"quantity" binding:"required"`
}
