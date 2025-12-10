package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.With(middlewares.JwtAuthMiddleware).Get("/auth/user/{id}", ur.userController.GetUserByIdHandler)	// protected route
	r.With(middlewares.CreateUserMiddleware).Post("/auth/signup", ur.userController.CreateUserHandler)
	r.With(middlewares.SignInUserMiddleware).Post("/auth/signin", ur.userController.SignInUserHandler)
}