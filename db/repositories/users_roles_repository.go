package db

import (
	"AuthInGo/models"
	"database/sql"
)

type IUsersRoles interface {
	GetUserRoles(userId int) ([]*models.Role, error)
	AssignRoleToUser(userId int, roleId int) (error)
	RemoveRoleFromUser(userId int, roleId int) (error)
	HasRole(userId int, roleName string) (bool, error)
	HasAllRoles(userId int, roleName []string) (bool, error)
	GetUserPermissions(userId int) ([]*models.Permission, error)
	HasPermission(userId int, permissionName string) (bool, error)
}

type UserRoles struct {
	db *sql.DB
}

func NewUserRoles(_db *sql.DB) (IUsersRoles) {
	return &UserRoles {
		db: _db,
	}
}

func (ur *UserRoles) GetUserRoles(userId int) ([]*models.Role, error) {

	var roles []*models.Role

	query := `
		SELECT
			ROLES.ID,
			ROLES.NAME,
			ROLES.DESCRIPTION,
			ROLES.CREATED_AT,
			ROLES.UPDATED_AT
		FROM ROLES
		INNER JOIN USERS_ROLES
			ON ROLES.ID = USERS_ROLES.ROLE_ID
		WHERE USERS_ROLES.USER_ID = ?
	`

	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role models.Role

		if err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.Created_at,
			&role.Updated_at,
		); err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}


func (ur *UserRoles) AssignRoleToUser(userId int, roleId int) error {

	query := `
		INSERT INTO USERS_ROLES (USER_ID, ROLE_ID) VALUES (?, ?)
	`

	_, err := ur.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}

	return nil
}


func (ur *UserRoles) RemoveRoleFromUser(userId int, roleId int) error {

	query := `
		DELETE FROM USERS_ROLES WHERE USER_ID = ? AND ROLE_ID = ?
	`

	_, err := ur.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}

	return nil
}


func (ur *UserRoles) HasRole(userId int, roleName string) (bool, error) {

	query := `
		SELECT ROLES.ID
		FROM ROLES
		INNER JOIN USERS_ROLES ON ROLES.ID = USERS_ROLES.ROLE_ID
		WHERE USERS_ROLES.USER_ID = ? AND ROLES.NAME = ?
		LIMIT 1
	`

	var roleId int
	err := ur.db.QueryRow(query, userId, roleName).Scan(&roleId)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (ur *UserRoles) HasAllRoles(userId int, roleNames []string) (bool, error) {

	query := `
		SELECT COUNT(*) = ? FROM USERS_ROLES INNER JOIN ROLES ON USERS_ROLES.ROLE_ID = ROLES.ID 
		WHERE USERS_ROLES.USER_ID = ? AND ROLES.NAME IN (?) GROUP BY USERS_ROLES.USER_ID
	`

	row := ur.db.QueryRow(query, len(roleNames), userId, roleNames)

	var hasAllRoles bool

	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}


	return hasAllRoles, nil
	
	// for _, role := range roleNames {
	// 	res, _ := ur.HasRole(userId, role)
	// 	if res == false {
	// 		return false, nil
	// 	}
	// }

	// return true, nil
	
}


func (ur *UserRoles) GetUserPermissions(userId int) ([]*models.Permission, error) {

	query := `
		SELECT DISTINCT
			PERMISSIONS.ID,
			PERMISSIONS.NAME,
			PERMISSIONS.DESCRIPTION,
			PERMISSIONS.RESOURCE,
			PERMISSIONS.ACTION,
			PERMISSIONS.CREATED_AT,
			PERMISSIONS.UPDATED_AT
		FROM USERS_ROLES
		INNER JOIN ROLES ON ROLES.ID = USERS_ROLES.ROLE_ID
		INNER JOIN ROLES_PERMISSIONS ON ROLES.ID = ROLES_PERMISSIONS.ROLE_ID
		INNER JOIN PERMISSIONS ON PERMISSIONS.ID = ROLES_PERMISSIONS.PERMISSION_ID
		WHERE USERS_ROLES.USER_ID = ?
	`

	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*models.Permission{}

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

func (ur *UserRoles) HasPermission(userId int, permissionName string) (bool, error) {

	query := `
		SELECT 1
		FROM USERS_ROLES
		INNER JOIN ROLES ON ROLES.ID = USERS_ROLES.ROLE_ID
		INNER JOIN ROLES_PERMISSIONS ON ROLES_PERMISSIONS.ROLE_ID = ROLES.ID
		INNER JOIN PERMISSIONS ON PERMISSIONS.ID = ROLES_PERMISSIONS.PERMISSION_ID
		WHERE USERS_ROLES.USER_ID = ? AND PERMISSIONS.NAME = ?
		LIMIT 1
	`

	var exists int
	err := ur.db.QueryRow(query, userId, permissionName).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
