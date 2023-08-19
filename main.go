package main

import (
	"github.com/Stellar-Lab/stellarminds-be/controllers"
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvironmentVariable()
	initializer.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("user/:id", middleware.RequireAuth, controllers.Show)

	r.Run()
}
