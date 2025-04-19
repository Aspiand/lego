package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name  string `json:"name"`
	Price uint   `json:"price"`

	ProductItems []ProductItems `gorm:"foreignKey:ProductID"` // 1:n
}

type ProductItems struct {
	gorm.Model

	ProductID string  `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}
