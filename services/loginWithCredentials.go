package services

import (
	"errors"
	"github.com/Stellar-Lab/stellarminds-be/initializer"
	"github.com/Stellar-Lab/stellarminds-be/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

func LoginWithCredentials(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
	}

	var user models.User
	if dbError := initializer.DB.First(&user, "email = ?", body.Email).Error; dbError != nil {
		if dbError != nil {
			if errors.Is(dbError, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Email doesn't exist!",
				})

				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Database Error: " + dbError.Error(),
			})

			return
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is Incorrect!",
		})

		return
	}

	// Generate JWT token. Set the exp date to 30 days.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the encoded token as a string using a jwt secret defined in .env
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate JWT token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		tokenString,
		3600*24*30,
		"",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, user)
}
