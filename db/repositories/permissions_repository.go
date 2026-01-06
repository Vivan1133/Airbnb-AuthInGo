package db

import (
	"AuthInGo/models"
	"database/sql"
)

type IPermissionsRepository interface {
	GetPermissionById(id int) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	DeletePermissionById(id int) (error)
	UpdatePermissionById(id int, name string, desc string, resource string, action string) (*models.Permission, error)
	CreatePermission(name string, desc string, resource string, action string) (*models.Permission, error)
}

type PermissionsRepository struct {
	db *sql.DB
}

// constructor
func NewPermissionsRepository(_db *sql.DB) IPermissionsRepository {
	return &PermissionsRepository {
		db: _db,
	}
}


func (r *PermissionsRepository) GetPermissionById(id int) (*models.Permission, error) {

	query := `
		SELECT ID, NAME, DESCRIPTION, RESOURCE, ACTION, CREATED_AT, UPDATED_AT
		FROM PERMISSIONS
		WHERE ID = ?
	`

	var p models.Permission
	err := r.db.QueryRow(query, id).Scan(
		&p.Id,
		&p.Name,
		&p.Description,
		&p.Resource,
		&p.Action,
		&p.Created_at,
		&p.Updated_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *PermissionsRepository) GetAllPermissions() ([]*models.Permission, error) {

	query := `
		SELECT ID, NAME, DESCRIPTION, RESOURCE, ACTION, CREATED_AT, UPDATED_AT
		FROM PERMISSIONS
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var p models.Permission

		if err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Resource,
			&p.Action,
			&p.Created_at,
			&p.Updated_at,
		); err != nil {
			return nil, err
		}

		permissions = append(permissions, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *PermissionsRepository) GetPermissionByName(name string) (*models.Permission, error) {

	query := `
		SELECT ID, NAME, DESCRIPTION, RESOURCE, ACTION, CREATED_AT, UPDATED_AT
		FROM PERMISSIONS
		WHERE NAME = ?
	`

	var p models.Permission
	err := r.db.QueryRow(query, name).Scan(
		&p.Id,
		&p.Name,
		&p.Description,
		&p.Resource,
		&p.Action,
		&p.Created_at,
		&p.Updated_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *PermissionsRepository) CreatePermission(name string, desc string, resource string, action string ) (*models.Permission, error) {

	query := `
		INSERT INTO PERMISSIONS (NAME, DESCRIPTION, RESOURCE, ACTION)
		VALUES (?, ?, ?, ?)
	`

	res, err := r.db.Exec(query, name, desc, resource, action)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetPermissionById(int(id))
}

func (r *PermissionsRepository) UpdatePermissionById(id int, name string, desc string, resource string, action string) (*models.Permission, error) {

	query := `
		UPDATE PERMISSIONS
		SET NAME = ?, DESCRIPTION = ?, RESOURCE = ?, ACTION = ?
		WHERE ID = ?
	`

	res, err := r.db.Exec(query, name, desc, resource, action, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return r.GetPermissionById(id)
}

func (r *PermissionsRepository) DeletePermissionById(id int) error {

	query := `
		DELETE FROM PERMISSIONS
		WHERE ID = ?
	`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
