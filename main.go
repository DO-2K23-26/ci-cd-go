package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// city represents data about a record city.
type City struct {
	gorm.Model
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Population float64 `json:"population"`
}

// cities slice to seed record city data.
var cities = []City{
	{ID: "1", Name: "Eaubonne", Country: "Rance", Population: 69420},
	{ID: "2", Name: "Cupacabana", Country: "Betweenmexicoandbuenosaires", Population: 8},
	{ID: "3", Name: "Worchestershire", Country: "Unitedkingdomland", Population: 49.9999},
}

/////// SHOULD BE REPLACED BY A CALL TO A BDD

func main() {
	// Get the environment variables
	apiAddr := os.Getenv("CITY_API_ADDR")
	apiPort := os.Getenv("CITY_API_PORT")

	dbURL := os.Getenv("CITY_API_DB_URL")
	dbUser := os.Getenv("CITY_API_DB_USER")
	dbPassword := os.Getenv("CITY_API_DB_PWD")
	dbName := os.Getenv("CITY_API_DB_NAME")

	dbconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", dbURL, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}

	// Migrate the schema
	db.AutoMigrate(&City{})

	// Insert the cities data in the database
	for _, city := range cities {
        err := db.Create(&city).Error
        if err != nil {
            fmt.Printf("Failed to insert city: %s\n", err)
        }
	}

	router := gin.Default()
	router.GET("/city", getCity)
	router.POST("/city", postCity)
	router.GET("/city/:id", getCityByID)
	router.GET("/_health", getHealth)

	err = router.Run("%s:%s", apiAddr, apiPort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %s", err))
	}
}

// getCity responds with the list of all cities as JSON.
func getCity(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cities)
}

// postCity adds an city from JSON received in the request body.
func postCity(c *gin.Context) {
	var newAlbum City

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new city to the slice.
	cities = append(cities, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getCityByID locates the city whose ID value matches the id
// parameter sent by the client, then returns that city as a response.
func getCityByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of cities, looking for
	// an city whose ID value matches the parameter.
	for _, a := range cities {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "city not found"})
}

// getHealth responds with the health of the service
func getHealth(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
