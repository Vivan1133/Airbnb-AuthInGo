package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

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
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/permissions", pr.permissionController.CreatePermission)
	// only admin can create a permission
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/permissions", pr.permissionController.GetAllPermissions)
	// only admin can get all permissions
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/permissions/{id}", pr.permissionController.GetPermissionById)
	// only admin can get permission by id
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Put("/permissions/{id}", pr.permissionController.UpdatePermission)
	// only admin can update a permission
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Delete("/permissions/{id}", pr.permissionController.DeletePermissionById)
	// only admin can delete a permission
}