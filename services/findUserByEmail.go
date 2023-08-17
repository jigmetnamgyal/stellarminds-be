package services

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindUserByEmail(c *gin.Context) (*models.User, error) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
	}

	var user models.User
	if dbError := initializer.DB.First(&user, "email = ?", body.Email).Error; dbError != nil {
		return nil, dbError
	}

	return &user, nil
}
