package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

type IUserService interface {
	GetUserById(id int64) (*models.User, error)
	Create(name string, email string, password string) (error)
	GetAllUser() ([]*models.User, error)
	DeleteUserById(id int64) error
	GetUserByEmail(email string) (*models.User, error)
	SignIn(email string, password string) (string, error)
}

type UserService struct {
	userRepository db.IUserRepository
	userRoleRepo   db.IUsersRoles
}

func NewUserService(_userRepository db.IUserRepository, userRoleRep db.IUsersRoles) IUserService {
	return &UserService{
		userRepository: _userRepository,
		userRoleRepo:userRoleRep,
	}
}

func (u *UserService) GetUserById(id int64) (*models.User, error) {

	user, err := u.userRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetAllUser() ([]*models.User, error) {

	users, err := u.userRepository.GetAll()

	if err != nil {
		fmt.Println("error getting users (in service)")
		return nil, err
	}

	return users, nil
}

func (u *UserService) DeleteUserById(id int64) error {

	err := u.userRepository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {

	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) Create(name string, email string, password string) (error) {

	encryptedPass := utils.HashPassword(password)

	lastInsertedId, err := u.userRepository.Create(name, email, encryptedPass)
	if err != nil {
		return err
	}
	// auto assign user role to all signed up user
	assignErr := u.userRoleRepo.AssignRoleToUser(lastInsertedId, 2)
	if assignErr != nil {
		return assignErr
	}
	return nil
}


func (u *UserService) SignIn (email string, password string) (string, error) {

	// fetch user by name
	// if not exists throw error
	// if exists match the plain password and hashed password
	// if matched signIn else throw error

	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("error while fetching user by email", err)
		return "", err
	}

	passwordMatch, passChechErr := utils.CheckHashedPassword(password, user.Password)

	if !passwordMatch {
		fmt.Println("password does not match")
		return "", passChechErr
	}

	token, tokenCreateErr := utils.CreateJwtToken(user.Email, user.Id)

	if tokenCreateErr != nil {
		fmt.Println("error while creating the jwt token")
		return "", tokenCreateErr
	}

	fmt.Println("successfully signed in")

	return token, nil

}