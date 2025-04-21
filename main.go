package main

import (
	"net/http"

	"github.com/Aspiand/lego/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routers.SetupRouters(router)

	router.Run("0.0.0.0:8000")
}
