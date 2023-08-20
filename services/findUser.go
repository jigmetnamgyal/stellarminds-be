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
	userId := c.Param("id")
	fmt.Println(userId)

	var user models.User
	if err := initializer.DB.Preload("Profile").First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User with id: " + userId + " is not found",
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
