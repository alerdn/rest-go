package models

import (
	"time"
)

type Sale struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nsu       int       `json:"nsu"`

	// belongsTo
	ProductId int     `json:"product_id"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductId"`

	// belongsTo
	UserId        int     `json:"user_id"`
	User          *User    `json:"user" gorm:"foreignKey:UserId"`
	Price         float64 `json:"price"`
	PaymentMethod string  `json:"payment_method"`
	Quantity      int     `json:"quantity"`
}
