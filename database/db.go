package database

import (
	"github.com/Aspiand/lego/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("database/lego.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Brand{}, &models.Product{})
}

func Get() *gorm.DB {
	if db == nil {
		Init()
	}

	return db
}
