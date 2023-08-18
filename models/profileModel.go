package models

import (
	"gorm.io/gorm"
	"time"
)

type GenderEnum string

const (
	GenderMale   GenderEnum = "male"
	GenderFemale GenderEnum = "female"
)

type Profile struct {
	gorm.Model
	UserId      uint       `json:"user_id" binding:"required"`
	Name        string     `gorm:"unique; not null" json:"name"  binding:"required"`
	DateOfBirth time.Time  `gorm:"not null" json:"date_of_birth"  binding:"required"`
	Gender      GenderEnum `gorm:"not null" json:"gender" binding:"required"`
	AvatarUrl   string     `gorm:"not null" json:"avatar_url" binding:"required"`
}
