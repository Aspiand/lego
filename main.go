package main

import (
	"net/http"

	"github.com/Aspiand/lego/database"
	"github.com/Aspiand/lego/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	routers.SetupRouters(router)
	database.Init()

	router.Run("localhost:8000")
}
