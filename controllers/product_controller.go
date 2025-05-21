package controllers

import (
	"net/http"

	"github.com/Aspiand/lego/internal/models"
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
	ctx := c.Request.Context()

	ctl.DB.WithContext(ctx).Preload("Brand").Find(&products)

	c.IndentedJSON(http.StatusOK, products)
}

func (ctl ProductController) GetByID(c *gin.Context) {
	var product models.Product

	if err := ctl.DB.Preload("Brand").First(&product, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func (ctl ProductController) Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.DB.Create(&product).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, product)
}

func (ctl ProductController) Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctl.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product)
	c.Status(http.StatusOK)
}

func (ctl ProductController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (ctl ProductController) DeleteAll(c *gin.Context) {
	result := ctl.DB.Where("1 = 1").Delete(&models.Product{})

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result.RowsAffected)
}
