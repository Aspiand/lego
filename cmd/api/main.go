package main

import (
	"github.com/Aspiand/lego/internal/database"
	"github.com/Aspiand/lego/internal/handler"
	"github.com/Aspiand/lego/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	dst := ":memory:"
	// dst = "database/lego.db"

	db := database.Initialize(dst)
	database.Seed(db)
	router := gin.Default()
	handler.RegisterBrandRoutes(router, handler.NewBrandHandler(service.NewBrandService(db)))
	router.Run("0.0.0.0:8000")
}
