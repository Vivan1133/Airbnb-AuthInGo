package db

import (
	// "database/sql"
	"fmt"
)

type IUserRepository interface {
	Create() error
}

type UserRepository struct {
	// db *sql.DB
}

func (u *UserRepository) Create() error {
	fmt.Println("Inside user repository Create()")
	return nil
}

func NewUserRepository() IUserRepository {
	return &UserRepository{

	}
}