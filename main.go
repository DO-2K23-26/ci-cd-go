package main

import (
	"encoding/json"

	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// The City struct represents data about a record city.
type City struct {
	ID             uint    `json:"id" gorm:"primaryKey;not null" `
	DepartmentCode string  `json:"department_code"`
	InseeCode      string  `json:"insee_code"`
	ZipCode        string  `json:"zip_code"`
	Name           string  `json:"name"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
}


func main() {
	// Get the environment variables
	apiAddr := os.Getenv("CITY_API_ADDR")
	apiPort := os.Getenv("CITY_API_PORT")

	dbURL := os.Getenv("CITY_API_DB_URL")
	dbUser := os.Getenv("CITY_API_DB_USER")
	dbPassword := os.Getenv("CITY_API_DB_PWD")
	dbName := os.Getenv("CITY_API_DB_NAME")

	path := "cities.json"

	dbconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbURL, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}

	// Migrate the schema
	err = db.AutoMigrate(&City{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate schema: %s", err))
	}

	// Seed the database
	err = seedData(path, db)
	if err != nil {
		fmt.Printf("failed to seed database: %s", err)
	}

	router := gin.Default()
	router.GET("/city", func(c *gin.Context) {
		err := getCity(db, c)
		if err != nil {  // Error handling
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get cities"})
		}
	})
	router.POST("/city", func(c *gin.Context) {
		err := postCity(db, c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to post"})
		}
	})
	router.GET("/city/:id", func(c *gin.Context) {
		err := getCityByID(db, c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get city"})
		}
	})
	router.GET("/_health", getHealth)

	err = router.Run(fmt.Sprintf("%s:%s", apiAddr, apiPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %s", err))
	}
}

func getCity(db *gorm.DB, c *gin.Context) error {
	var cities []City
	err := db.Find(&cities).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get cities"})
		return err
	}
	c.IndentedJSON(http.StatusOK, cities)
	return nil
}

// postCity adds an city from JSON received in the request body.
func postCity(db *gorm.DB, c *gin.Context) error {
	var newCity City
	err := c.BindJSON(&newCity)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return err
	}

	err = db.Create(&newCity).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to create city"})
		return err
	}

	c.IndentedJSON(http.StatusCreated, newCity)
	return nil
}

// getCityByID locates the city whose ID value matches the id
// parameter sent by the client, then returns that city as a response.
func getCityByID(db *gorm.DB, c *gin.Context) error {
	var city City
	id := c.Param("id")
	result := db.First(&city, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "city not found"})
		return result.Error
	}
	c.IndentedJSON(http.StatusOK, city)
	return nil
}

func seedData(path string, db *gorm.DB) error {
	// Open the cities JSON file, handle errors
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return err
	}
	defer file.Close()

	// Decode the JSON cities file into an array
	decoder := json.NewDecoder(file)
	var cities []City
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

// getHealth responds with the health of the service
func getHealth(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
