package controllers

import (
	"github.com/Stellar-Lab/stellarminds-be/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	services.UserCreator(c)
}

func LogIn(c *gin.Context) {
	services.LoginWithCredentials(c)
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in!",
	})
}

func Show(c *gin.Context) {
	services.FindUser(c)
}

func Index(c *gin.Context) {
	services.QueryAllUsers(c)
}
