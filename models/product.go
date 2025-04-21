package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey" `
	Name      string `gorm:"size:255;not null"`
	Price     uint   `gorm:"index;not null;default:0"`
	BrandID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Brand Brand         // belongs to relationship
	Items []ProductItem // one to many relationship
}

type ProductItem struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	SelledAt  time.Time

	Product Product // belongs to relationship
}
