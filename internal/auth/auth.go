package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAuthToken(userID uint) (string, error) {
	if userID == 0 {
		return "", errors.New("userID cannot be zero")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	jwtkey := os.Getenv("JWT_SECRET_KEY")
	if jwtkey == "" {
		return "", errors.New("JWT_SECRET_KEY is not set")
	}

	return token.SignedString([]byte(jwtkey))
}
