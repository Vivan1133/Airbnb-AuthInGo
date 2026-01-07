package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.IRoleService
}

func NewRoleController(roleService services.IRoleService) (*RoleController) {
	return &RoleController {
		roleService: roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	roleIdString := chi.URLParam(r, "roleId")

	roleId, err := strconv.Atoi(roleIdString)
	if err != nil {
		http.Error(w, "wrong user id passed", http.StatusBadRequest)
		return
	}

	role, err := rc.roleService.GetRoleById(roleId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to fetch role", err)
		return
	}

	if role == nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "role not found", fmt.Errorf("role with id %v not found", roleId))
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Role fetched successfully", role)

}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.roleService.GetAllRoles()
	if err != nil {
		fmt.Println(err)
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to fetch roles", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "fetched all roles SUCCESSFULLY", roles)
}