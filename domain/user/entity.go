package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	ID          int8      `json:"id"`
	Name        string    `json:"name" gorm:"size:255" gorm:"not null"`
	Email       string    `json:"email" gorm:"size:100"`
	Password    string    `json:"password" gorm:"size:100"`
	PhoneNumber string    `json:"phone_number" gorm:"size:50"`
	Role        string    `json:"role" gorm:"size:50"`
	Gender      int8      `json:"gender"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
