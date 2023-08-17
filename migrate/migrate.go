package main

import (
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"log"
)

func init() {
	initializer.LoadEnvironmentVariable()
	initializer.ConnectToDb()
}

func main() {
	err := initializer.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error migrating")
	}
}
