package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string  `gorm:"unique; not null" validate:"email" json:"email" binding:"required"`
	Password        string  `gorm:"not null" json:"password" binding:"required"`
	ConfirmPassword string  `gorm:"not null" json:"confirm_password"`
	AgreeToTerms    bool    `json:"agree_to_terms"`
	Profile         Profile `json:"profile"`
}
