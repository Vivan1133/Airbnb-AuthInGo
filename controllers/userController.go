package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(_userService services.IUserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside user controller")
	uc.userService.GetUserById()
	w.Write([]byte("user reg endpoint"))
}