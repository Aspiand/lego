package database

import (
	"github.com/Aspiand/lego/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize(destination string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(destination), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database")
	}

	db.AutoMigrate(
		&models.Brand{},
		&models.Product{},
		&models.ProductItem{},
	)

	return db
}

func Seed(db *gorm.DB) {
	db.Create([]*models.Brand{
		{Name: "Intel"},
		{Name: "AMD"},
		{Name: "Axioo"},
		{Name: "Gigabyte"},
		{Name: "Acome"},
	})

	db.Create([]*models.Product{
		{
			Name:    "Core 2 Duo E8400",
			Price:   12000,
			BrandID: 1,
		},

		{
			Name:    "Axioo Hype 5 Gen 12 Core I5 1235u 16GB",
			Price:   6000000,
			BrandID: 3,
		},

		{
			Name:    "Motherboard G41 DDR 3",
			Price:   150000,
			BrandID: 4,
		},

		{
			Name:  "Monitor",
			Price: 350000,
			Brand: &models.Brand{
				Name: "Zyrexx",
			},
		},

		{
			Name:  "Ram 4 + 2 GB DDR3",
			Price: 120000,
			Brand: nil,
		},
	})
}
