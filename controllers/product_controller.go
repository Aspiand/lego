package controllers

import (
	"net/http"

	"github.com/Aspiand/lego/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (ctl ProductController) GetAll(ctx *gin.Context) {
	var products []models.Product
	ctl.DB.Preload("Items").Find(&products)
	ctx.IndentedJSON(http.StatusOK, products)
}
