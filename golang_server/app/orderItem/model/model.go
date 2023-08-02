package model

type OrderItem struct {
	ID        int `json:"id" db:"id"`
	Quantity  int `json:"quantity" db:"quantity"`
	OrderID   int `json:"order_id" db:"order_id"`
	ProductID int `json:"product_id" db:"product_id"`
}

func (OrderItem) TableName() string { return "order_item" }
