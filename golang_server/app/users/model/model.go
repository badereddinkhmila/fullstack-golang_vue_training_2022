package model

import (
	"time"
)

type User struct {
	Id        int       `json:"id"         db:"id"`
	Username  string    `json:"username"   db:"username"`
	Email     string    `json:"email"      db:"email"`
	Password  string    `json:"password"   db:"password"`
	Role      string    `json:"role" 	   db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (User) TableName() string { return "app_users" }

const (
	ROLE_SUPERADMIN string = "ROLE_SUPERADMIN"
	ROLE_PROVIDER   string = "ROLE_PROVIDER"
	ROLE_CUSTOMER   string = "ROLE_CUSTOMER"
)
