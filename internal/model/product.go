package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text;not null"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Quantity    int     `json:"quantity" gorm:"type:int;not null"`
}
