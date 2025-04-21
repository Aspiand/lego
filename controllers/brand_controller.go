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

func (ctl BrandController) GetAll(ctx *gin.Context) {
	var brands []models.Brand
	ctl.DB.Find(&brands)

	ctx.IndentedJSON(http.StatusOK, brands)
}

func (ctl BrandController) Create(ctx *gin.Context) { // test later
	var brand models.Brand

	if err := ctx.ShouldBindJSON(&brand); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.DB.Create(&brand).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, brand)
}

func (ctl BrandController) Delete(ctx *gin.Context) { // later
	id := ctx.Param("id")

	if err := ctl.DB.Delete(&models.Brand{}, id).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
