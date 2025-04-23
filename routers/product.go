package routers

import (
	"github.com/Aspiand/lego/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerProductRoutes(db *gorm.DB, rg *gin.Engine) {
	product := rg.Group("/products")
	controller := controllers.NewProductController(db)
	{
		product.GET("/", controller.GetAll)
		product.DELETE("/", controller.DeleteAll)
	}
}
