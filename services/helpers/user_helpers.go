package helpers

import "github.com/Stellar-Lab/stellarminds-be/models"

var validGenders = map[models.GenderEnum]bool{
	models.GenderMale:   true,
	models.GenderFemale: true,
	models.GenderOther:  true,
}

func ValidGender(input models.GenderEnum) bool {
	return validGenders[input]
}
