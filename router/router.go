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

func CreateRouter(UserRouter Router, RoleRouter Router, PermissionRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiterMiddleware)

	chiRouter.Get("/ping", controllers.PingHandler)

	chiRouter.HandleFunc("/fakestore/*", utils.ProxyToService("https://fakestoreapi.com/", "/fakestore"))
	// http://localhost:3004/fakestore/users/1				// https://fakestoreapi.com/	// /fakestore
															// https://fakestoreapi.com/users/1


	chiRouter.HandleFunc("/hotelservice/*", utils.ProxyToService("http://localhost:3000/", "/hotelservice"))														
	// http://localhost:3004/hotelservice/api/v1/hotels/3		// http://localhost:3000/	// /hotelservice
																// http://localhost:3000/api/v1/hotels/3


	chiRouter.HandleFunc("/bookingservice/*", utils.ProxyToService("http://localhost:3002/", "/bookingservice"))
	// http:localhost:3004/bookingservice/*						// http://localhost:3002/*		
	
	chiRouter.HandleFunc("/reviewservice/*", utils.ProxyToService("http://localhost:8081/", "/reviewservice"))
	// http:localhost:3004/reviewservice/*						// http://localhost:8081/*

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)
	PermissionRouter.Register(chiRouter)

	return chiRouter

}