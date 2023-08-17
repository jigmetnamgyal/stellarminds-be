package initializer

import (
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email           string
		Password        string
		ConfirmPassword string
		AgreeToTerms    bool
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password doesn't match",
		})
	}

	hash, hashPasswordError := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if hashPasswordError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while hashing the password",
		})

		return
	}

	user := models.User{
		Email:           body.Email,
		Password:        string(hash),
		ConfirmPassword: string(hash),
		AgreeToTerms:    body.AgreeToTerms,
	}

	result := DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
