package models

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ExpiryDate  time.Time `json:"expiryDate"`
	CategoryID  uint      `json:"categoryId"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	ImagePath   string    `json:"imagePath"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
