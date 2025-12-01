package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type IUserService interface {
	CreateUser() error
}

type UserService struct {
	userRepository db.IUserRepository
}

func NewUserService(_userRepository db.IUserRepository) IUserService {
	return &UserService{
		userRepository: _userRepository,
	}
}

func (u *UserService) CreateUser() error {
	fmt.Println("Inside user service createUser")
	u.userRepository.Create()
	return nil
}
