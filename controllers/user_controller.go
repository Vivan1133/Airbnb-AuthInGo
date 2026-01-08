package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/middlewares"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
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

	userIdString := chi.URLParam(r, "id")	// returned string id

	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		http.Error(w, "wrong user id passed", http.StatusBadRequest)
		return
	}

	user, err := uc.userService.GetUserById(int64(userID))

	if err != nil {
		http.Error(w, "could not find the user", http.StatusInternalServerError)
		return
	}

	errJson := utils.WriteJsonSucessResponse(w, http.StatusOK, "user found", user)

	if errJson != nil {
		http.Error(w, "error writing json success response", http.StatusInternalServerError)
	}

	// uc.userService.GetUserById()
	// w.Write([]byte("user reg endpoint"))
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	payloadRaw := r.Context().Value(middlewares.CreateUserCtxKey)

	fmt.Println("Raw payload: ", payloadRaw)

	payload := payloadRaw.(dtos.CreateUserRequestDto)

	err := uc.userService.Create(payload.Name, payload.Email, payload.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while creating user", err)
		return
	}

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