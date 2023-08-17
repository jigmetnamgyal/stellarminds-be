package services

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UserCreator(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})

		return
	}

	validate := validator2.New()

	if validationError := validate.Struct(body); validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not valid",
		})

		return
	}

	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm password doesn't match your password",
		})

		return
	}

	hash, hashError := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if hashError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Email:           body.Email,
		Password:        string(hash),
		ConfirmPassword: string(hash),
		AgreeToTerms:    body.AgreeToTerms,
	}

	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
