package main

import (
	"net/http"

	"example.com/turbosteel_api/controllers"
	"example.com/turbosteel_api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/alldogs", controllers.AllDogs)

	router.Run()
}
