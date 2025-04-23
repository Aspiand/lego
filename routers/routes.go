package routers

import (
	"github.com/Aspiand/lego/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouters(db *gorm.DB, router *gin.Engine) {
	registerProductRoutes(db, router)
	registerBrandRoutes(db, router)

	router.GET("/test", func(ctx *gin.Context) {
		// db.Create(&models.Brand{Name: "Intel"})
		db.Create(&models.Product{Name: "Core 2 Duo E8400", Price: 12000, BrandID: 1})
	})
}
