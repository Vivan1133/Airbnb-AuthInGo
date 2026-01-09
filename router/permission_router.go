package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type PermissionRouter struct {
	permissionController *controllers.PermissionController
}

func NewPermissionRouter(permissionController *controllers.PermissionController) Router {
	return &PermissionRouter{
		permissionController: permissionController,
	}
}

func (pr *PermissionRouter) Register(r chi.Router) {
	r.Post("/permissions", pr.permissionController.CreatePermission)
	r.Get("/permissions", pr.permissionController.GetAllPermissions)
	r.Get("/permissions/{id}", pr.permissionController.GetPermissionById)
	r.Put("/permissions/{id}", pr.permissionController.UpdatePermission)
	r.Delete("/permissions/{id}", pr.permissionController.DeletePermissionById)
}