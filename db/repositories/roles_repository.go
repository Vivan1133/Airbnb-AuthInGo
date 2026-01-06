package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type IRoleRepository interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, desc string) (error)
	DeleteRoleById(id int) (error)
	UpdateRoleById(id int, name string, desc string) (error)
}

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) IRoleRepository {
	return &RoleRepository {
		db: _db,
	}
}

func (rr *RoleRepository) GetRoleById(id int) (*models.Role, error) {

	var role models.Role

	query := `SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT FROM ROLES WHERE ID = ?`
	row := rr.db.QueryRow(query, id)

	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found with ID: ", id);
			return nil, err
		}
		return nil, err
	}

	return &role, nil
}

func (rr *RoleRepository) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role

	query := `SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT FROM ROLE WHERE NAME = ?`
	row := rr.db.QueryRow(query, name)

	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found with NAME: ", name)
			return nil, err
		}
		return nil, err
	}

	return &role, nil

}

func (rr *RoleRepository) GetAllRoles() ([]*models.Role, error) {
	var roles []*models.Role

	query := `SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT FROM ROLES`
	rows, err := rr.db.Query(query)
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Empty Roles tables", err)
			return nil, err
		}
		return nil, err
	}

	for rows.Next() {
		var role models.Role

		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.Created_at, &role.Updated_at); err != nil {
			fmt.Println("Error while scanning roles row ERROR: ", err)
		}
		roles = append(roles, &role)

	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

	return roles, nil
}

func (rr *RoleRepository) CreateRole(name string, desc string) (error) {

	query := `INSERT INTO ROLES (NAME, DESCRIPTION) VALUES (?, ?)`
	result, err := rr.db.Exec(query, name, desc)

	if err != nil {
		fmt.Println("Can not insert into roles ERROR: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("Error in rowsAffected")
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("0 rows affected")
		return nil
	}

	fmt.Println("Successfully INSERTED in ROLES")
	return nil
}

func (rr *RoleRepository) DeleteRoleById(id int) (error) {

	query := `DELETE FROM ROLES WHERE ID = ?`

	result, err := rr.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error while executing query ERROR: ", err)
		return nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error in rows affected ERROR: ", err)
		return nil
	}
	if rowsAffected == 0 {
		fmt.Println("0 rows affected")
		return nil
	}

	fmt.Println("Successfully DELETED the record");
	return nil

}

func (rr *RoleRepository) UpdateRoleById(id int, name string, desc string) (error) {

	query := `UPDATE ROLES SET NAME = ?, DESCRIPTION = ? WHERE ID = ?`

	result, err := rr.db.Exec(query, name, desc, id)

	if err != nil {
		fmt.Println("Failed to execute update query ERROR: ", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Println("Rows affected ERROR: ", err)
		return nil
	}

	if rowsAffected == 0 {
		fmt.Println("0 rows affected")
		return nil
	}

	fmt.Println("Successfully updated the Roles record")

	return nil
}