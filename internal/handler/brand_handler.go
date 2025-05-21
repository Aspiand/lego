package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Aspiand/lego/internal/config"
	"github.com/Aspiand/lego/internal/models"
	"github.com/Aspiand/lego/internal/service"
	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	service *service.BrandService
}

func NewBrandHandler(service *service.BrandService) *BrandHandler {
	return &BrandHandler{
		service: service,
	}
}

func RegisterBrandRoutes(r *gin.Engine, h *BrandHandler) {
	group := r.Group("/brands")
	group.GET("", h.ListBrands)
	group.POST("", h.Create)
	group.GET("/:id", h.GetById)
	group.PATCH("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

func (h *BrandHandler) ListBrands(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil && page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.DefaultPageSize)))
	if err != nil && pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	brandName := c.DefaultQuery("name", "")

	brands, err := h.service.List(ctx, brandName, page, pageSize)

	if err != nil {
		// if errors.Is(err, context.DeadlineExceeded) {}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brands"})
		return
	}

	c.JSON(http.StatusOK, brands)
}

func (h *BrandHandler) GetById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
	}

	brand, err := h.service.GetByID(ctx, brandID)
	if err != nil {
	}

	c.JSON(http.StatusOK, brand)
}

func (h *BrandHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	var brand models.Brand

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(ctx, &brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
	}

	c.JSON(http.StatusCreated, brand)
}

func (h *BrandHandler) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil && brandID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}

	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(ctx, brandID, brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *BrandHandler) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), config.DefaultTimeout)
	defer cancel()

	brandID, err := strconv.Atoi(c.Param("id"))
	if err != nil && brandID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.service.Delete(ctx, brandID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
