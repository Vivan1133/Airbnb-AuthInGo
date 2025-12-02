package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type IUserService interface {
	GetUserById() error
}

type UserService struct {
	userRepository db.IUserRepository
}

func NewUserService(_userRepository db.IUserRepository) IUserService {
	return &UserService{
		userRepository: _userRepository,
	}
}

func (u *UserService) GetUserById() error {
	fmt.Println("Inside user service")
	u.userRepository.DeleteById(1)
	return nil
}
