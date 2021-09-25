package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string) (bool, uint64) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, 0
	}
	claims := token.Claims.(jwt.MapClaims)
	uid := uint64(claims["user_id"].(float64))
	return true, uid
}
