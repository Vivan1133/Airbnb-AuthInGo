package db

import (
	"AuthInGo/models"
	"database/sql"
)

type IRolesPermissions interface {
	AssignPermissionToRole(roleId int, permissionId int) (error)
	RemovePermissionFromRole(roleId int, permissionId int) (error)
	RoleHasPermission(roleId int, permissionName string) (bool, error)
	GetRolePermissions(roleId int) ([]*models.Permission, error) 
}

type RolesPermissions struct {
	db *sql.DB
}

func NewRolesPermissions(_db *sql.DB) (IRolesPermissions) {
	return &RolesPermissions {
		db: _db,
	}
}

func (rp *RolesPermissions) AssignPermissionToRole(roleId int, permissionId int) error {

	query := `
		INSERT INTO ROLES_PERMISSIONS (ROLE_ID, PERMISSION_ID)
		VALUES (?, ?)
	`

	_, err := rp.db.Exec(query, roleId, permissionId)
	if err != nil {
		return err
	}

	return nil
}

func (rp *RolesPermissions) RemovePermissionFromRole(roleId int, permissionId int) error {

	query := `
		DELETE FROM ROLES_PERMISSIONS
		WHERE ROLE_ID = ? AND PERMISSION_ID = ?
	`

	_, err := rp.db.Exec(query, roleId, permissionId)
	if err != nil {
		return err
	}

	return nil
}


func (rp *RolesPermissions) RoleHasPermission(roleId int, permissionName string) (bool, error) {

	query := `
		SELECT 1
		FROM ROLES_PERMISSIONS
		INNER JOIN PERMISSIONS
			ON PERMISSIONS.ID = ROLES_PERMISSIONS.PERMISSION_ID
		WHERE ROLES_PERMISSIONS.ROLE_ID = ?
		  AND PERMISSIONS.NAME = ?
		LIMIT 1
	`

	var exists int
	err := rp.db.QueryRow(query, roleId, permissionName).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (rp *RolesPermissions) GetRolePermissions(roleId int) ([]*models.Permission, error) {

	query := `
		SELECT
			PERMISSIONS.ID,
			PERMISSIONS.NAME,
			PERMISSIONS.DESCRIPTION,
			PERMISSIONS.RESOURCE,
			PERMISSIONS.ACTION,
			PERMISSIONS.CREATED_AT,
			PERMISSIONS.UPDATED_AT
		FROM ROLES_PERMISSIONS
		INNER JOIN PERMISSIONS
			ON PERMISSIONS.ID = ROLES_PERMISSIONS.PERMISSION_ID
		WHERE ROLES_PERMISSIONS.ROLE_ID = ?
	`

	rows, err := rp.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var p models.Permission

		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Resource,
			&p.Action,
			&p.Created_at,
			&p.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}


