package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJSONToken(tokenData map[string]interface{}, jwtSecret string, expiresAt int) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Second) * time.Duration(expiresAt)).Unix()
	claims := jwt.MapClaims{}

	claims["expiredAt"] = expiredAt
	claims["authorization"] = true

	for i, v := range tokenData {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := to.SignedString([]byte(jwtSecret))

	if err != nil {
		return token, err
	}
	return token, nil
}
