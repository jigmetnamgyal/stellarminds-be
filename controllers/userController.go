package controllers

import (
	"errors"
	"github.com/Stellar-Lab/stellarminds-be/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SignUp(c *gin.Context) {
	services.UserCreator(c)
}

func LogIn(c *gin.Context) {
	user, err := services.FindUserByEmail(c)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "record not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Database Error: " + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
