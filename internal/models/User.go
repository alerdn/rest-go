package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`

	// hasMany
	Sales *[]Sale `json:"sales" gorm:"foreignKey:UserId"`
}
