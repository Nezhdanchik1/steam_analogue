package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"week_9_crud/internal/db"
	"week_9_crud/internal/models"
)

var secret = []byte("Nezhdanchik1")

func GenerateJWT(userId int) (string, error) {
	var user models.User
	db.DB.First(&user, userId)

	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateJWT(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("invalid claims")
	}

	return token, claims, nil
}
