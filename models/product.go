package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name    string `gorm:"size:255;not null"`
	Price   uint   `gorm:"index;not null;default:0"`
	BrandID uint

	Brand Brand `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items []ProductItem
}

type ProductItem struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	SelledAt  gorm.DeletedAt

	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
