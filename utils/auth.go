package utils

import (
	env "AuthInGo/config/env"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainTextPassword string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error encrypting password", err)
	}

	return string(hashed)
}

func CheckHashedPassword(plainTextPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))

	if err != nil {
		fmt.Println("password matching failed")
		return false, err
	} else {
		fmt.Println("password matched")
		return true, nil
	}
}

func CreateJwtToken(email string) (string, error) {

	secretKey := []byte(env.GetString("SECRET_KEY", "mysecretkey"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
