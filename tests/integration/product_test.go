package integration_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"testing"

// 	"github.com/Aspiand/lego/internal/models"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetProduct(t *testing.T) {
// 	db, router := SetupTest()

// 	var product = models.Product{Name: "Air Doa", Price: 100000}
// 	db.Create(&product)

// 	req, _ := http.NewRequest("GET", "/products/1", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var rProduct models.Product
// 	json.Unmarshal(w.Body.Bytes(), &rProduct)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, product.Name, rProduct.Name)
// 	assert.Nil(t, rProduct.Brand)
// }

// func TestCreateProduct(t *testing.T) {
// 	db, router := SetupTest()

// 	var product = models.Product{Name: "Air Minum Bekas ...", Price: 5000}
// 	payload := toJSONReader(product)

// 	req, _ := http.NewRequest("POST", "/products", payload)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var dProduct models.Product
// 	db.First(&dProduct, 1)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Equal(t, product.Name, dProduct.Name)
// 	assert.Equal(t, product.Price, dProduct.Price)
// }

// func TestUpdateProduct(t *testing.T) {
// 	db, router := SetupTest()
// 	db.Create(&models.Product{Name: "Garam", Price: 5000})

// 	updatedProduct := models.Product{Name: "Garam Rukyah", Price: 150000}
// 	payload := toJSONReader(updatedProduct)

// 	req, _ := http.NewRequest("PUT", "/products/1", payload)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var product models.Product
// 	db.First(&product, 1)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, updatedProduct.Name, product.Name)
// 	assert.Equal(t, updatedProduct.Price, product.Price)
// }

// func TestDeleteProduct(t *testing.T) {
// 	db, router := SetupTest()

// 	db.Create(&models.Brand{Name: "Kopi Bekas ..."})

// 	req, _ := http.NewRequest("DELETE", "/brands/1", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var product models.Brand
// 	db.Unscoped().First(&product, 1)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.False(t, product.DeletedAt.Time.IsZero())
// }

// func TestDeleteAll(t *testing.T) {
// 	db, router := SetupTest()

// 	var products = []models.Product{
// 		{Name: "Air Doa"},
// 		{Name: "Garam Rukyah"},
// 		{Name: "Posisi"},
// 	}
// 	db.Create(products)

// 	req, _ := http.NewRequest("DELETE", "/products", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, strconv.Itoa(len(products)), w.Body.String())
// }
