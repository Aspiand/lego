package routers

import (
	"github.com/Aspiand/lego/controllers"
	"github.com/Aspiand/lego/database"
	"github.com/gin-gonic/gin"
)

func registerBrandRoutes(rg *gin.Engine) {
	brand := rg.Group("/brands")
	controller := controllers.NewBrandController(database.Get())
	{
		brand.GET("/", controller.GetAll)
		brand.DELETE("/:id", controller.Delete)
	}
}
