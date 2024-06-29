package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"cicd-go/models"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.City{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func TestGetCities(t *testing.T) {
	db := setupTestDB()
	router := setupRouter()

	// Insert test data
	db.Create(&models.City{Name: "Test City"})

	router.GET("/city", func(c *gin.Context) {
		GetCities(c, db)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/city", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var cities []models.City
	err := json.Unmarshal(w.Body.Bytes(), &cities)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cities))
	assert.Equal(t, "Test City", cities[0].Name)
}

func TestPostCity(t *testing.T) {
	db := setupTestDB()
	router := setupRouter()

	router.POST("/city", func(c *gin.Context) {
		PostCity(c, db)
	})

	newCity := models.City{Name: "Test City"}
	jsonValue, _ := json.Marshal(newCity)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/city", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var createdCity models.City
	err := json.Unmarshal(w.Body.Bytes(), &createdCity)
	assert.NoError(t, err)
	assert.Equal(t, "Test City", createdCity.Name)

	var cities []models.City
	db.Find(&cities)
	assert.Equal(t, 2, len(cities))
	assert.Equal(t, "Test City", cities[0].Name)
}

func TestGetCityByID(t *testing.T) {
	db := setupTestDB()
	router := setupRouter()

	// Insert test data
	city := models.City{Name: "Test City"}
	db.Create(&city)

	router.GET("/city/:id", func(c *gin.Context) {
		GetCityByID(c, db)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/city/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var returnedCity models.City
	err := json.Unmarshal(w.Body.Bytes(), &returnedCity)
	assert.NoError(t, err)
	assert.Equal(t, "Test City", returnedCity.Name)
}
