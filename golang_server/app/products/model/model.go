package model

import "time"

type Product struct {
	ID                 int       `json:"id"                  db:"id"`
	NameProduct        string    `json:"name_product"        db:"name_product"`
	DescriptionProduct string    `json:"description_product" db:"description_product"`
	Price              float64   `json:"price_product"       db:"price_product"`
	Quantity           int       `json:"quantity_product"    db:"quantity"`
	IsAvailable        bool      `json:"is_available"        db:"is_available"`
	ProductImage       []byte    `json:"product_image"       db:"image_product"`
	CreatedAt          time.Time `json:"created_at"          db:"created_at"`
	CategoryID         int       `json:"category_id"         db:"category_id"`
	CreatedBy          int       `json:"provider_id"         db:"created_by"`
}

func (Product) TableName() string { return "products" }
