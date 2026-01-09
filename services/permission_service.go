package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type IPermissionService interface {
	CreatePermission(name string, desc string, resource string, action string) (*models.Permission, error)

	GetPermissionById(id int) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)

	UpdatePermission(id int, name string, desc string, resource string, action string) (*models.Permission, error)

	DeletePermissionById(id int) error
}

type PermissionService struct {
	permissionRepository db.IPermissionsRepository
}

func NewPermissionsService(permissionRepo db.IPermissionsRepository) (IPermissionService) {
	return &PermissionService {
		permissionRepository: permissionRepo,
	}
}

func (ps *PermissionService) CreatePermission(name, desc, resource, action string ) (*models.Permission, error) {
	return ps.permissionRepository.CreatePermission(
		name, desc, resource, action,
	)
}

func (ps *PermissionService) GetPermissionById(id int) (*models.Permission, error) {
	return ps.permissionRepository.GetPermissionById(id)
}

func (ps *PermissionService) GetPermissionByName(name string) (*models.Permission, error) {
	return ps.permissionRepository.GetPermissionByName(name)
}

func (ps *PermissionService) GetAllPermissions() ([]*models.Permission, error) {
	return ps.permissionRepository.GetAllPermissions()
}

func (ps *PermissionService) UpdatePermission(id int, name, desc, resource, action string) (*models.Permission, error) {
	return ps.permissionRepository.UpdatePermissionById(
		id, name, desc, resource, action,
	)
}

func (ps *PermissionService) DeletePermissionById(id int) error {
	return ps.permissionRepository.DeletePermissionById(id)
}

