package db

import (
	// "database/sql"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type IUserRepository interface {
	Create() error
	GetById() (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
}

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) GetAll() ([]*models.User, error) {

	query := "SELECT * FROM USERS"

	rows, err := u.db.Query(query)

	if err != nil {
		fmt.Println("error getting all data")
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}

		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)

		if err != nil {
			fmt.Println("error while scanning")
			return nil, err
		}

		users = append(users, user)
	}

	rowsErr := rows.Err()

	if rowsErr != nil {
		fmt.Println("rows error")
		return nil, rowsErr
	}

	fmt.Println("successfully fetched all users")

	for index, value := range users {
		fmt.Println(index, value)
	}

	return users, nil
}

func (u *UserRepository) DeleteById(id int64) error {

	query := "DELETE FROM USERS WHERE id = ?"

	res , err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("error executing delete query", err)
		return err
	}

	rowsAffected, rowsErr := res.RowsAffected()

	if rowsErr != nil {
		fmt.Println("rows error occurred")
		return rowsErr
	}

	if rowsAffected == 0 {
		fmt.Println("no rows affected");
		return nil
	}

	fmt.Println("successfully deleted ")

	return nil
}

func(u *UserRepository) Create() error {

	// query
	query := "INSERT INTO USERS (name, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, "vansh", "vansh@sample.com", "654321")

	if err != nil {
		fmt.Println("error inserting the user")
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("error fetching rows affected")
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("no rows were affected")
		return nil
	}

	fmt.Println("successfully created the user, rows affected: ", rowsAffected)

	return nil
}

func (u *UserRepository) GetById() (*models.User, error) {
	fmt.Println("Inside user repository")

	// prepare the query

	query := "SELECT ID, NAME, EMAIL, CREATED_AT, UPDATED_AT FROM USERS WHERE ID = ?"

	row := u.db.QueryRow(query, 1)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No such user found with the given id")
			return nil, err
        } else {
			fmt.Println("error scanning user", err)
			return nil, err
		}
	}

	fmt.Println("printing user: ", user)
	return user, err
}

func NewUserRepository(_db *sql.DB) IUserRepository {
	return &UserRepository{
		db: _db,
	}
}