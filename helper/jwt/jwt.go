package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(payload map[string]interface{}, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	for key, value := range payload {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(expiry).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	coba, err := token.SignedString([]byte{})
	return coba, err
}
