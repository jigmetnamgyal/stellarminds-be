package models

import (
	"gorm.io/gorm"
	"time"
)

type GenderEnum string

const (
	GenderMale   GenderEnum = "male"
	GenderFemale GenderEnum = "female"
	GenderOther  GenderEnum = "others"
)

type Profile struct {
	gorm.Model
	UserId      uint       `json:"-"`
	Name        string     `gorm:"not null" json:"name"`
	DateOfBirth time.Time  `gorm:"not null" json:"date_of_birth"`
	Gender      GenderEnum `gorm:"not null" json:"gender"`
	AvatarUrl   string     `gorm:"not null" json:"avatar_url"`
}
