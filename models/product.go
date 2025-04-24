package models

import (
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name    string        `json:"name" gorm:"size:255;not null"`
	Price   uint          `json:"price" gorm:"index;not null;default:0"`
	BrandID uint          `json:"brand_id" gorm:"default:0"`
	Brand   *Brand        `json:"brand" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items   []ProductItem `json:"-"`
}

type ProductItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"index;not null"`
	Product   Product        `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SelledAt  gorm.DeletedAt `json:"selled_at"`
}
