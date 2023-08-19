package services

import (
	"errors"
	"fmt"
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func FindUser(c *gin.Context) {
	userId := c.Query("id")
	fmt.Println(userId)

	var user models.User
	if err := initializer.DB.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User with id: " + userId + " is not found",
			})

			c.AbortWithStatus(http.StatusNotFound)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, user)
}
