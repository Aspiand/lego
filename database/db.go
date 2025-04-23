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
