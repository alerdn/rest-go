package sale

import "time"

type Sale struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Nsu       int       `json:"nsu"`
	ProductId int       `json:"product_id"`
	UserId    int       `json:"user_id"`
	Price     float64   `json:"price"`
}
