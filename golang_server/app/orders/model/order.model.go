package model

import "time"

type Order struct {
	ID          int       `json:"id"         db:"id"`
	UserID      int       `json:"user_id"    db:"user_id"`
	OrderStatus string    `json:"status"     db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func (Order) TableName() string { return "orders" }
