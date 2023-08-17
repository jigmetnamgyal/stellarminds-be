package controllers

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
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

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
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

	c.JSON(http.StatusOK, gin.H{})
}
