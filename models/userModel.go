package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string `gorm:"unique; not null" validate:"email"`
	Password        string
	ConfirmPassword string
	AgreeToTerms    bool
}
