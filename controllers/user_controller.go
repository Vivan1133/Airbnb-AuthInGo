package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/middlewares"
	"AuthInGo/services"
	"AuthInGo/utils"
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




	// uc.userService.GetUserById()
	w.Write([]byte("user reg endpoint"))
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	payloadRaw := r.Context().Value(middlewares.CreateUserCtxKey)

	fmt.Println("Raw payload: ", payloadRaw)

	payload := payloadRaw.(dtos.CreateUserRequestDto)

	uc.userService.Create(payload.Name, payload.Email, payload.Password)

	utils.WriteJsonSucessResponse(w, http.StatusOK, "successfully created the user", "")

}

func (uc *UserController) SignInUserHandler(w http.ResponseWriter, r *http.Request) {

	payloadRaw := r.Context().Value(middlewares.SignInUserCtx)

	payload := payloadRaw.(dtos.SignInUserRequestDto)

	token, err := uc.userService.SignIn(payload.Email, payload.Password)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "error while sigin process", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "successfully signed in", token)
	
}