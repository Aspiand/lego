package routers

import (
	"github.com/Aspiand/lego/controllers"
	"github.com/Aspiand/lego/database"
	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.Engine) {
	product := rg.Group("/products")
	controller := controllers.NewProductController(database.Get())
	{
		product.GET("/", controller.GetAll)
	}
}
