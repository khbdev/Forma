package token

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAdminToken() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expiresStr := os.Getenv("JWT_EXPIRES_MINUTES")

	expiresMinutes, err := strconv.Atoi(expiresStr)
	if err != nil || expiresMinutes <= 0 {
		expiresMinutes = 30
	}

	claims := jwt.MapClaims{
		"sub":  "admin",
		"role": "admin",
		"exp":  time.Now().Add(time.Duration(expiresMinutes) * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}