package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenJWT string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func DecodeToken(tokenJWT string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenJWT)

	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid Token")
}
