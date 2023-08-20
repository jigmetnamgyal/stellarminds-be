package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutUser(c *gin.Context) {
	c.SetCookie(
		"Authorization",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "You are successfully Logged Out"})
}
