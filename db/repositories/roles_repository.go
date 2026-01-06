package db

import (
	"AuthInGo/models"
	"database/sql"
)

type IRoleRepository interface {

	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)

	CreateRole(name string, desc string) (*models.Role, error)
	UpdateRoleById(id int, name string, desc string) (*models.Role, error)

	DeleteRoleById(id int) error

	RoleExistsById(id int) (bool, error)
	RoleExistsByName(name string) (bool, error)
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

	query := `
		SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT
		FROM ROLES
		WHERE ID = ?
	`

	var role models.Role
	err := rr.db.QueryRow(query, id).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.Created_at,
		&role.Updated_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}


func (rr *RoleRepository) GetRoleByName(name string) (*models.Role, error) {

	query := `
		SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT
		FROM ROLES
		WHERE NAME = ?
	`

	var role models.Role
	err := rr.db.QueryRow(query, name).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.Created_at,
		&role.Updated_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &role, nil
}


func (rr *RoleRepository) GetAllRoles() ([]*models.Role, error) {

	query := `
		SELECT ID, NAME, DESCRIPTION, CREATED_AT, UPDATED_AT
		FROM ROLES
	`

	rows, err := rr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role

	for rows.Next() {
		var role models.Role
		err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.Created_at,
			&role.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}


func (rr *RoleRepository) CreateRole(name string, desc string) (*models.Role, error) {

	query := `
		INSERT INTO ROLES (NAME, DESCRIPTION)
		VALUES (?, ?)
	`

	result, err := rr.db.Exec(query, name, desc)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return rr.GetRoleById(int(id))
}


func (rr *RoleRepository) DeleteRoleById(id int) error {

	query := `
		DELETE FROM ROLES WHERE ID = ?
	`

	_, err := rr.db.Exec(query, id)
	return err
}


func (rr *RoleRepository) UpdateRoleById(id int, name string, desc string) (*models.Role, error) {

	query := `
		UPDATE ROLES
		SET NAME = ?, DESCRIPTION = ?
		WHERE ID = ?
	`

	_, err := rr.db.Exec(query, name, desc, id)
	if err != nil {
		return nil, err
	}

	return rr.GetRoleById(id)
}

func (rr *RoleRepository) RoleExistsById(id int) (bool, error) {

	query := `
		SELECT 1 FROM ROLES WHERE ID = ? LIMIT 1
	`

	var exists int
	err := rr.db.QueryRow(query, id).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (rr *RoleRepository) RoleExistsByName(name string) (bool, error) {

	query := `
		SELECT 1 FROM ROLES WHERE NAME = ? LIMIT 1
	`

	var exists int
	err := rr.db.QueryRow(query, name).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
