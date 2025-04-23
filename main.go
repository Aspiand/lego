package main

import (
	"github.com/Aspiand/lego/database"
	"github.com/Aspiand/lego/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	dst := "file::memory:?cache=shared"
	// dst = "database/lego.db"

	db := database.Initialize(dst)
	router := gin.Default()
	routers.SetupRouters(db, router)
	router.Run("0.0.0.0:8000")
}
