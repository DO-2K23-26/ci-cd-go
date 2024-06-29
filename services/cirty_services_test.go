package services

import (
	"cicd-go/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

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

	// Insert test data
	db.Create(&models.City{Name: "Test City"})

	cities, err := GetCities(db)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(cities))
	assert.Equal(t, "Test City", cities[0].Name)
}

func TestGetCityByID(t *testing.T) {
	db := setupTestDB()

	// Insert test data
	city := models.City{Name: "Test City"}
	db.Create(&city)
	str := strconv.Itoa(int(city.ID))
	returnedCity, err := GetCityByID(db, str)
	assert.NoError(t, err)
	assert.Equal(t, "Test City", returnedCity.Name)
}

func TestAddCity(t *testing.T) {
	db := setupTestDB()

	newCity := models.City{Name: "Test City"}
	err := AddCity(db, newCity)
	assert.NoError(t, err)

	var cities []models.City
	db.Find(&cities)
	assert.Equal(t, 3, len(cities))
	assert.Equal(t, "Test City", cities[0].Name)
}
