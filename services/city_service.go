package services

import (
	"encoding/json"
	"fmt"
	"os"
	"gorm.io/gorm"
	"cicd-go/models"
)

func GetCities(db *gorm.DB) ([]models.City, error) {
	var cities []models.City
	err := db.Find(&cities).Error
	return cities, err
}

func GetCityByID(db *gorm.DB, id string) (*models.City, error) {
	var city models.City
	err := db.First(&city, id).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func AddCity(db *gorm.DB, newCity models.City) error {
	err := db.Create(&newCity).Error
	return err
}

func SeedData(path string, db *gorm.DB) error {
	// Open the cities JSON file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return err
	}
	defer file.Close()

	// Decode the JSON cities file into an array
	decoder := json.NewDecoder(file)
	var cities []models.City
	err = decoder.Decode(&cities)
	if err != nil {
		fmt.Printf("Failed to decode JSON: %s\n", err)
		return err
	}

	// Seed (write) array to database
	for _, city := range cities {
		err := db.Create(&city).Error
		if err != nil {
			fmt.Printf("Failed to insert city: %s\n", err)
		}
	}
	return nil
}
