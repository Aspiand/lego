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
		product.GET("", controller.GetAll)
		product.POST("", controller.Create)
		product.DELETE("", controller.DeleteAll)
		product.GET("/:id", controller.GetByID)
		product.PUT("/:id", controller.Update)
		product.DELETE("/:id", controller.Delete)
	}
}
