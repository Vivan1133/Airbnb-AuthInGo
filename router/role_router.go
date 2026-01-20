package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(roleController *controllers.RoleController) Router {
	return &RoleRouter {
		roleController: roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/roles/id/{roleId}", rr.roleController.GetRoleById)
	// only admin can get role by id
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/roles", rr.roleController.GetAllRoles)
	// only admin can get all roles
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/roles/name/{roleName}", rr.roleController.GetRoleByName)
	// only admin can get role by name
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin"), middlewares.CreateRoleMiddleware).Post("/roles", rr.roleController.CreateRole)
	// only admin can create a role
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Delete("/roles/id/{roleId}", rr.roleController.DeleteRole)
	// only admin can delete a role
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin"), middlewares.UpdateRoleMiddleware).Patch("/roles", rr.roleController.UpdateRole)
	// only admin can update a role
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/roles-permissions/{roleId}/{permissionId}", rr.roleController.AssignPermission)
	// only admin can assign permission to a role
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Delete("/roles-permissions/{roleId}/{permissionId}", rr.roleController.RemovePermission)
	// only admin can remove permission from a role
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/roles-permissions/{roleId}", rr.roleController.GetRolePermissions)
	// only admin can get role permissions
	r.With(middlewares.JwtAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/users-roles/assign/{userId}/{roleId}", rr.roleController.AssignRoleToUser)
	// to assign a role only admin is allowed
}