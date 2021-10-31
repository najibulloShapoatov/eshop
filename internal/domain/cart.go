package domain

import "time"

type Cart struct {
	ID        int64      `json:"id" db:"id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	Items     []*CartItem `json:"items" db:"-"`
}

func NewCart(userId int64) *Cart {
	now := time.Now()
	return &Cart{
		UserID:    userId,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
