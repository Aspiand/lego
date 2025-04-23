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

func (ctl ProductController) GetAll(c *gin.Context) {
	var products []models.Product
	ctl.DB.Preload("Items").Find(&products)
	c.IndentedJSON(http.StatusOK, products)
}

func (ctl ProductController) DeleteAll(c *gin.Context) {
	if err := ctl.DB.Where("1 = 1").Delete(&models.Product{}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
