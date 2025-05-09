package routers

import (
	"github.com/Aspiand/lego/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerBrandRoutes(db *gorm.DB, rg *gin.Engine) {
	brand := rg.Group("/brands")
	controller := controllers.NewBrandController(db)
	{
		brand.GET("", controller.GetAll)
		brand.POST("", controller.Create)
		brand.DELETE("", controller.DeleteAll)
		brand.GET("/:id", controller.GetByID)
		brand.PUT("/:id", controller.Update)
		brand.DELETE("/:id", controller.Delete)
	}
}
