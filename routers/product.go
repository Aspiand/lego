package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.Engine) {
	product := rg.Group("/products")
	{
		product.GET("/", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"message": "ehe",
			})
		})
	}
}
