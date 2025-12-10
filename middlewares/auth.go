package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type userIDCtx struct{}
var UserIDCtx = userIDCtx{}
type userEmailCtx struct{}
var UserEmailCtx = userEmailCtx{}

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "auth header required", http.StatusUnauthorized)
			return 
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "auth Header must have a Bearer", http.StatusUnauthorized)
			return 
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "empty token provided", http.StatusUnauthorized)
			return 
		}

		claims := jwt.MapClaims{}

		t, err := jwt.ParseWithClaims(token, &claims, func (token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("SECRET_KEY", "secretkey")), nil
		})

		if err != nil {
			fmt.Println(t, err)
			http.Error(w, "invalid token passed", http.StatusUnauthorized)
			return 
		}

		email, ok := claims["email"].(string)
		id, idOk := claims["id"]

		if !ok || !idOk {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return 
		}

		fmt.Printf("authentication successfull user with email as %v and id as %v", email, id)
		
		ctx := context.WithValue(r.Context(), UserIDCtx, id)
		ctx = context.WithValue(ctx, UserEmailCtx, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}