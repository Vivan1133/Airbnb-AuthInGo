package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
)

//
type Router interface {
	Register(r chi.Router)
}

func CreateRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiterMiddleware)

	chiRouter.Get("/ping", controllers.PingHandler)

	chiRouter.HandleFunc("/fakestore/*", utils.ProxyToService("https://fakestoreapi.com/", "/fakestore"))
	// http://localhost:3001/fakestore/users/1				// https://fakestoreapi.com/	// /fakestore
															// https://fakestoreapi.com/users/1



	UserRouter.Register(chiRouter)

	return chiRouter

}