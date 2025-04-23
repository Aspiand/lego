package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Aspiand/lego/models"
	"github.com/stretchr/testify/assert"
)

func TestGetBrand(t *testing.T) {
	db, router := SetupTest()

	var brand = models.Brand{Name: "AMD"}
	db.Create(&brand)

	req, _ := http.NewRequest("GET", "/brands/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var rBrand models.Brand
	json.Unmarshal(w.Body.Bytes(), &rBrand)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, brand.Name, rBrand.Name)
}

func TestCreateBrand(t *testing.T) {
	_, router := SetupTest()
	brandName := "Intel"

	payload := toJSONReader(models.Brand{Name: brandName})
	req, _ := http.NewRequest("POST", "/brands", payload)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var brand models.Brand
	json.Unmarshal(w.Body.Bytes(), &brand)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, brandName, brand.Name)
}

func TestUpdateBrand(t *testing.T) {
	db, router := SetupTest()

	db.Create(&models.Brand{Name: "ehe"})

	payload := toJSONReader(models.Brand{Name: "EHE"})
	req, _ := http.NewRequest("PUT", "/brands/1", payload)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var brand models.Brand
	db.First(&brand, 1)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "EHE", brand.Name)
}

func TestDeleteBrand(t *testing.T) {
	db, router := SetupTest()

	db.Create(&models.Brand{Name: "Axioo"})

	req, _ := http.NewRequest("DELETE", "/brands/1", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var brand models.Brand
	db.Unscoped().First(&brand, 1)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.False(t, brand.DeletedAt.Time.IsZero())
}

func TestDeleteAllBrand(t *testing.T) {
	db, router := SetupTest()

	var brands = []*models.Brand{
		{Name: "AMD"},
		{Name: "Intel"},
		{Name: "Axioo"},
	}
	db.Create(brands)

	req, _ := http.NewRequest("DELETE", "/brands", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strconv.Itoa(len(brands)), w.Body.String())
}
