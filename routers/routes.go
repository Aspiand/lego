package routers

import (
	"net/http"

	"github.com/Aspiand/lego/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouters(db *gorm.DB, router *gin.Engine) {
	registerProductRoutes(db, router)
	registerBrandRoutes(db, router)

	if true {
		router.GET("/reset", func(ctx *gin.Context) {
			db.Where("1 = 1").Delete(&models.Brand{})
			db.Where("1 = 1").Delete(&models.Product{})
			ctx.Redirect(http.StatusFound, "/products")
		})
	}
}
