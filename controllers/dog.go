// controllers/dog.go

package controllers

import (
	"net/http"

	"example.com/turbosteel_api/models"
	"github.com/gin-gonic/gin"
)

func AllDogs(context *gin.Context) {
	var dogs []models.Dog
	models.DB.Find(&dogs)

	context.JSON(http.StatusOK, gin.H{"data": dogs})
}
