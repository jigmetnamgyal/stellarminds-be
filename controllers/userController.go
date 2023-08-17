package controllers

import (
	"github.com/Stellar-Lab/stellarminds-be/services"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	services.UserCreator(c)
}

func LogIn(c *gin.Context) {
	user, err := services.FindUserByEmail(c)

	if err != nil {
		return
	}

	//passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	//
	//if passwordError != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Incorrect Password!",
	//	})
	//}

	//c.JSON(http.StatusOK, user)
}
