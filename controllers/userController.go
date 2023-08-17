package controllers

import (
	"github.com/Stellar-Lab/stellarminds-be/services"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	services.UserCreator(c)
}
