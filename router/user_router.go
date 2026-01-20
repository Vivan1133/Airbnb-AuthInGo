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
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).Get("/auth/user/{id}", ur.userController.GetUserByIdHandler)
	// someone whos trying to get user by id must at least be a user
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).Get("/auth/users", ur.userController.GetAllUsersHandler)
	// someone whos trying to get all users must at least be a user
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Delete("/auth/user/{id}", ur.userController.DeleteUserByIdHandler)
	// only admin can delete a user
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).Get("/auth/user/email/{email}", ur.userController.GetUserByEmailHandler)
	// someone whos trying to get user by email must at least be a user
	r.With(middlewares.CreateUserMiddleware).Post("/auth/signup", ur.userController.CreateUserHandler)
	r.With(middlewares.SignInUserMiddleware).Post("/auth/signin", ur.userController.SignInUserHandler)
}