package initializer

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(DB)
	if err != nil {
		log.Fatal("Error connecting to db")
	}
}
