package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string `gorm:"unique; not null"`
	Password        string
	ConfirmPassword string
	AgreeToTerms    bool
}
