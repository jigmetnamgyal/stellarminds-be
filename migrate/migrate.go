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
	initializer.DB.Model(&models.User{}).Association("Profile")

	err := initializer.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error migrating User")
	}

	profileMigrationError := initializer.DB.AutoMigrate(&models.Profile{})

	if profileMigrationError != nil {
		log.Fatal("Error migrating profile")
	}
}
