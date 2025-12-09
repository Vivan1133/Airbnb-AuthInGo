package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

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

	UserRouter.Register(chiRouter)

	return chiRouter

}