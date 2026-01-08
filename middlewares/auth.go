package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strings"
	repo "AuthInGo/db/repositories"
	dbConfig "AuthInGo/config/db"
	"github.com/golang-jwt/jwt/v5"
)

type userIDCtx struct{}
var UserIDCtx = userIDCtx{}
type userEmailCtx struct{}
var UserEmailCtx = userEmailCtx{}

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "auth header required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "auth header must start with Bearer", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			http.Error(w, "empty token provided", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("SECRET_KEY", "secretkey")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// --- Email ---
		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "invalid email claim", http.StatusUnauthorized)
			return
		}

		// --- User ID  ---
		idFloat, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "invalid user id claim", http.StatusUnauthorized)
			return
		}

		userID := int(idFloat)

		fmt.Printf(
			"authentication successful: email=%s, id=%d\n",
			email, userID,
		)

		ctx := context.WithValue(r.Context(), UserIDCtx, userID)
		ctx = context.WithValue(ctx, UserEmailCtx, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func RequireAllRoles(roles ...string) func(http.Handler) (http.Handler) {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userId := r.Context().Value(UserIDCtx).(int)

			db , _ := dbConfig.SetupDB()

			urr := repo.NewUserRoles(db)

			hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId, roles)

			if hasAllRolesErr != nil {
				fmt.Println(hasAllRolesErr)
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return 
			}

			fmt.Println("User got all roles")

			next.ServeHTTP(w, r)
		})
	}

}

func RequireAnyRole(roles ...string) func(http.Handler) (http.Handler) {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value(UserIDCtx).(int)

			db , _ := dbConfig.SetupDB()

			urr := repo.NewUserRoles(db)

			hasAnyRole, hasAnyRoleErr := urr.HasAnyRole(userId, roles)

			if hasAnyRoleErr != nil {
				fmt.Println(hasAnyRoleErr)
				http.Error(w, "Error checking user roles", http.StatusInternalServerError)
				return
			}

			if !hasAnyRole {
				http.Error(w, "Forbidden: You do not have any of the required roles", http.StatusForbidden)
				return 
			}

			fmt.Println("User got role")

			h.ServeHTTP(w, r)
		})
	}
}