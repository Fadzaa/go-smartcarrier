package course

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	ID          int8      `json:"id" gorm:"primaryKey"`
	ImageCourse string    `json:"image_course"`
	Title       string    `json:"title" gorm:"type:varchar(100)"`
	Description string    `json:"description"`
	Percentage  int8      `json:"percentage"`
	Difficulty  string    `json:"difficulty" gorm:"type:varchar(20)"`
	Category    string    `json:"category" gorm:"type:varchar(20)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
