package main

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvironmentVariable()
	initializer.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
