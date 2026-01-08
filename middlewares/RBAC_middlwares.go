package middlewares

import (
	"AuthInGo/dtos"
	"AuthInGo/utils"
	"context"
	"net/http"
)

type createRoleCtxKey struct{}
var CreateRoleCtxKey = createRoleCtxKey{}

type updateRoleCtxKey struct{}
var UpdateRoleCtxKey = updateRoleCtxKey{}


func CreateRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var payload dtos.CreateRoleDTO

		jsonReadErr := utils.ReadJsonRequest(r, &payload)
		if jsonReadErr != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while reading create role request", jsonReadErr)
			return 
		}

		ctx := context.WithValue(r.Context(), CreateRoleCtxKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))


	})
}

func UpdateRoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.UpdateRoleDTO

		jsonReadErr := utils.ReadJsonRequest(r, &payload)

		if jsonReadErr != nil {
			utils.WriteErrorResponse(w, http.StatusForbidden, "error while reading update req data", jsonReadErr)
			return 
		}

		ctx := context.WithValue(r.Context(), UpdateRoleCtxKey, &payload)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}