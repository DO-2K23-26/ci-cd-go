package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"cicd-go/models"
	"cicd-go/services"
)

func GetCities(c *gin.Context, db *gorm.DB) {
	cities, err := services.GetCities(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get cities"})
		return
	}
	c.IndentedJSON(http.StatusOK, cities)
}

func GetCityByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	city, err := services.GetCityByID(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "city not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, city)
}

func PostCity(c *gin.Context, db *gorm.DB) {
	var newCity models.City
	if err := c.BindJSON(&newCity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := services.AddCity(db, newCity)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to create city"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCity)
}
