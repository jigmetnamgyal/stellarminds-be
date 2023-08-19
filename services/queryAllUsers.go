package services

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAllUsers(c *gin.Context) {
	var users []models.User
	result := initializer.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, users)
}
