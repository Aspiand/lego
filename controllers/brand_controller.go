package controllers

import (
	"net/http"

	"github.com/Aspiand/lego/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandController struct {
	DB *gorm.DB
}

func NewBrandController(db *gorm.DB) *BrandController {
	return &BrandController{DB: db}
}

func (ctl BrandController) GetAll(c *gin.Context) {
	var brands []models.Brand
	ctx := c.Request.Context()

	if err := ctl.DB.WithContext(ctx).Find(&brands).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, brands)
}

func (ctl BrandController) GetByID(c *gin.Context) {
	var brand models.Brand

	if err := ctl.DB.First(&brand, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, brand)
}

func (ctl BrandController) Create(c *gin.Context) {
	var brand models.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.DB.Create(&brand).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, brand)
}

func (ctl BrandController) Update(c *gin.Context) {
	var brand models.Brand
	id := c.Param("id")

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctl.DB.Model(&models.Brand{}).Where("id = ?", id).Updates(brand)
	c.Status(http.StatusOK)
}

func (ctl BrandController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.DB.Delete(&models.Brand{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (ctl BrandController) DeleteAll(c *gin.Context) {
	result := ctl.DB.Where("1 == 1").Delete(&models.Brand{})

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result.RowsAffected)
}
