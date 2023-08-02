package model

import "time"

type Category struct {
	ID          int       `json:"id"          db:"id"`
	Lable       string    `json:"lable"       db:"lable"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at"  db:"created_at"`
	CreatedBy   int       `json:"created_by"  db:"created_by"`
}

func (Category) TableName() string { return "categories" }
