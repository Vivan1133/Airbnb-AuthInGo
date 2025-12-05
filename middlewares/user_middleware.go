package middlewares

import (
	"AuthInGo/dtos"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

type createUserCtxKey struct{}
var CreateUserCtxKey = createUserCtxKey{}

func CreateUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		fmt.Println("inside create user middleware")

		var payload dtos.CreateUserRequestDto

		jsonErr := utils.ReadJsonRequest(r, &payload)

		if jsonErr != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "failed to read request body", jsonErr)
		}

		validErr := utils.Validator.Struct(&payload)

		if validErr != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "validation failed", validErr)
			return
		}

		ctx := context.WithValue(r.Context(), CreateUserCtxKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type signInUserCtx struct{}
var SignInUserCtx = signInUserCtx{}

func SignInUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var payload dtos.SignInUserRequestDto

		if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "error while reading request", jsonErr)
		}

		if validatonErr := utils.Validator.Struct(&payload); validatonErr != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid req body", validatonErr)
		}

		ctx := context.WithValue(r.Context(), SignInUserCtx, payload)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}