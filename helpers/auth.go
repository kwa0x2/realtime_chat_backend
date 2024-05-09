package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

func GenerateToken(id string, email string, photo string) (string, error) {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"photo": photo,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func GetClaims(tokenString string) (map[string]interface{}, error) {
	err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	_, parseErr := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if parseErr != nil{
		return nil, parseErr
	}


	return claims, nil
}