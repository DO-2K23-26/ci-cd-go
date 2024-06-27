package routes

import (
	"cicd-go/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCityRoutes(router *gin.Engine, db *gorm.DB) {
	cityRoutes := router.Group("/city")
	{
		cityRoutes.GET("/", func(c *gin.Context) { controllers.GetCities(c, db) })
		cityRoutes.GET("/:id", func(c *gin.Context) { controllers.GetCityByID(c, db) })
		cityRoutes.POST("/", func(c *gin.Context) { controllers.PostCity(c, db) })
	}
}
