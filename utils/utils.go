package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const(
	tokenExpiry = time.Minute * 15
)

func GenerateJWT(userID string, email string)(string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"user_id":userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(tokenExpiry).Unix(),
	}
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := tokenObject.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		return "", err
	}

	return signedToken, nil
}