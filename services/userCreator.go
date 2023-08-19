package services

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/Stellar-Lab/stellarminds-be/services/helpers"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UserCreator(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body for user" + err.Error(),
		})

		return
	}

	var gender models.GenderEnum
	gender = body.Profile.Gender

	if !helpers.ValidGender(gender) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Gender!",
		})

		return
	}

	validate := validator2.New()

	if validationError := validate.Struct(body); validationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is not valid" + validationError.Error(),
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
		Profile: models.Profile{
			Name:        body.Profile.Name,
			DateOfBirth: body.Profile.DateOfBirth,
			Gender:      body.Profile.Gender,
			AvatarUrl:   body.Profile.AvatarUrl,
		},
	}

	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User" + result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
