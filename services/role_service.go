package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"database/sql"
	"errors"
)

type IRoleService interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)

	CreateRole(name string, desc string) (*models.Role, error)
	UpdateRole(id int, name string, desc string) (*models.Role, error)
	DeleteRole(id int) error

	RoleExists(id int) (bool, error)
	RoleNameAvailable(name string) (bool, error)

	AssignPermission(roleId int, permissionId int) error
	RemovePermission(roleId int, permissionId int) error
	GetRolePermissions(roleId int) ([]*models.Permission, error)

}

type RoleService struct {
	roleRepository db.IRoleRepository
	rolePermRepo   db.IRolesPermissions
	permissionRepo db.IPermissionsRepository
}

func NewRoleService(roleRepo db.IRoleRepository, rolePermRepo db.IRolesPermissions, permissionRepo db.IPermissionsRepository) (IRoleService) {
	return &RoleService {
		roleRepository: roleRepo,
		rolePermRepo: rolePermRepo,
		permissionRepo: permissionRepo, 
	}
}

func (rs *RoleService) GetRoleById(id int) (*models.Role, error) {
	return rs.roleRepository.GetRoleById(id)
}

func (rs *RoleService) GetRoleByName(name string) (*models.Role, error) {
	return rs.roleRepository.GetRoleByName(name)
}

func (rs *RoleService) GetAllRoles() ([]*models.Role, error) {
	return rs.roleRepository.GetAllRoles()
}

func (rs *RoleService) CreateRole(name string, desc string) (*models.Role, error) {

	exists, error := rs.roleRepository.RoleExistsByName(name)

	if error != nil {
		return nil, error
	}

	if exists {
		return nil, errors.New("Role already exists")
	}

	return rs.roleRepository.CreateRole(name, desc)

}

func (rs *RoleService) UpdateRole(id int, name string, desc string) (*models.Role, error) {

	exists, err := rs.roleRepository.RoleExistsById(id)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, sql.ErrNoRows
	}

	return rs.roleRepository.UpdateRoleById(id, name, desc)
}

func (rs *RoleService) DeleteRole(id int) error {

	exists, err := rs.roleRepository.RoleExistsById(id)
	if err != nil {
		return err
	}

	if !exists {
		return sql.ErrNoRows
	}

	return rs.roleRepository.DeleteRoleById(id)
}

func (rs *RoleService) RoleExists(id int) (bool, error) {
	return rs.roleRepository.RoleExistsById(id)
}

func (rs *RoleService) RoleNameAvailable(name string) (bool, error) {
	exists, err := rs.roleRepository.RoleExistsByName(name)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (rs *RoleService) AssignPermission(roleId int, permissionId int) (error) {

	roleExists, err := rs.roleRepository.RoleExistsById(roleId)
	if err != nil {
		return err
	}
	if !roleExists {
		return errors.New("role does not exist")
	}

	perm, err := rs.permissionRepo.GetPermissionById(permissionId)
	if err != nil {
		return err
	}
	if perm == nil {
		return errors.New("permission does not exist")
	}

	return rs.rolePermRepo.AssignPermissionToRole(roleId, permissionId)
}

func (rs *RoleService) RemovePermission(roleId int, permissionId int) error {
	return rs.rolePermRepo.RemovePermissionFromRole(roleId, permissionId)
}

func (rs *RoleService) GetRolePermissions(roleId int) ([]*models.Permission, error) {

	exists, err := rs.RoleExists(roleId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, sql.ErrNoRows
	}

	return rs.rolePermRepo.GetRolePermissions(roleId)
}
