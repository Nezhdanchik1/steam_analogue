package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secret = []byte("Nezhdanchik1")

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
