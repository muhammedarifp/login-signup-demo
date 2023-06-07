package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("jwt-123")

func CreateJwtToken(name string, email string) (string, error) {
	clims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clims)
	SignedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return SignedToken, nil
}

// func CheckValidJwtToken() bool {
// 	jwt.
// }
