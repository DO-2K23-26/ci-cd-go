package main

import (
	"fmt"
	"net/http"
	"os"

	"cicd-go/models"
	"cicd-go/routes"
	"cicd-go/services"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	err = db.AutoMigrate(&models.City{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate schema: %s", err))
	}

	// Seed the database
	err = services.SeedData("cities.json", db)
	if err != nil {
		fmt.Printf("failed to seed database: %s", err)
	}

	router := gin.Default()
	routes.RegisterCityRoutes(router, db)
	router.GET("/_health", getHealth)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	err = router.Run(fmt.Sprintf("%s:%s", apiAddr, apiPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %s", err))
	}
}

// getHealth responds with the health of the service
func getHealth(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
