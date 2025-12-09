package utils

import (
	env "AuthInGo/config/env"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
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

type visitor struct {
	limiter *rate.Limiter
	lastSeen time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

func GetVisitor(ip string) *rate.Limiter {

	defer mu.Unlock()
	
	mu.Lock()
	v, exists := visitors[ip]

	if !exists {
		limiter := rate.NewLimiter(rate.Every(1 * time.Minute), 5)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter

}