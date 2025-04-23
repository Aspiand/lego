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

func (ctl BrandController) Delete(c *gin.Context) { // later
	id := c.Param("id")

	if err := ctl.DB.Delete(&models.Brand{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
