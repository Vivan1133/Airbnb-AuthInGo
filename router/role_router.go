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
	r.Get("/roles/{roleId}", rr.roleController.GetRoleById)
	r.Get("/roles", rr.roleController.GetAllRoles)
	r.Get("/roles/{roleName}", rr.roleController.GetRoleByName)
	r.With(middlewares.CreateRoleMiddleware).Post("/roles", rr.roleController.CreateRole)
	r.Delete("/roles/{roleId}", rr.roleController.DeleteRole)
	r.With(middlewares.UpdateRoleMiddleware).Patch("/roles", rr.roleController.UpdateRole)
	r.Post("/roles-permissions/{roleId}/{permissionId}", rr.roleController.AssignPermission)
	r.Delete("/roles-permissions/{roleId}/{permissionId}", rr.roleController.RemovePermission)
	r.Get("/roles-permissions/{roleId}", rr.roleController.GetRolePermissions)
}