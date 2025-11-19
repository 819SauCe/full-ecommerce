package helpers

import (
	"errors"
	"full-ecommerce/internal/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv(config.Jwt_secret))

type UserData struct {
	Id          string
	Profile_img string
	First_name  string
	Last_name   string
	Email       string
	Role        string
}

func GenerateToken(user UserData) (string, error) {
	claims := jwt.MapClaims{
		"sub":         user.Id,
		"profile_img": user.Profile_img,
		"first_name":  user.First_name,
		"last_name":   user.Last_name,
		"email":       user.Email,
		"role":        user.Role,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
