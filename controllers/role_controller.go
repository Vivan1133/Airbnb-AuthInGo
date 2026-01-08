package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/middlewares"
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


func (rc *RoleController) GetRoleByName(w http.ResponseWriter, r *http.Request) {
	// 1. retrieve name from request param
	// 2. call service function
	// 3. retrieve record in a variable
	// 4. send json res back

	roleName := chi.URLParam(r, "roleName")

	role, roleErr := rc.roleService.GetRoleByName(roleName)

	if roleErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "error occured while fetching role", roleErr)
		return
	}

	if role == nil {
		utils.WriteJsonResponse(w, http.StatusNotFound, nil)
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Role fetched successfully", role)

}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {

	payloadRaw := r.Context().Value(middlewares.CreateRoleCtxKey)
	payload := payloadRaw.(dtos.CreateRoleDTO)

	role, createRoleErr := rc.roleService.CreateRole(payload.Name, payload.Description)

	if createRoleErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while creating role", createRoleErr)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusAccepted, "Successfully created the role", role)

}

func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := chi.URLParam(r, "roleId")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid role ID sent", err)
		return
	}

	deleteErr := rc.roleService.DeleteRole(roleId)
	if deleteErr != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error occured while deleting role", deleteErr)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Successfully deleted the role", nil)

}

func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {

	payloadRaw := r.Context().Value(middlewares.UpdateRoleCtxKey)

	payload := payloadRaw.(*dtos.UpdateRoleDTO)

	userId, convErr := strconv.Atoi(payload.Id)
	if convErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Wrong userID passed", convErr)
		return
	}

	role, err := rc.roleService.UpdateRole(userId, payload.Name, payload.Description)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error occured while updating the role", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Successfully updated the role", role)

}

func (rc *RoleController) AssignPermission(w http.ResponseWriter, r *http.Request) {

	roleIdString := chi.URLParam(r, "roleId")
	permissionIdStr := chi.URLParam(r, "permissionId")

	roleId, roleErr := strconv.Atoi(roleIdString)
	if roleErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid roleId passed", roleErr)
		return
	}
	permId, permErr := strconv.Atoi(permissionIdStr)
	if permErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid permission Id passed", permErr)
		return
	}

	err := rc.roleService.AssignPermission(roleId, permId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while assigning perm", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Successfully assigned permission to role", nil)

}

func (rc *RoleController) RemovePermission(w http.ResponseWriter, r *http.Request) {
	roleIdString := chi.URLParam(r, "roleId")
	permissionIdStr := chi.URLParam(r, "permissionId")

	roleId, roleErr := strconv.Atoi(roleIdString)
	if roleErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid roleId passed", roleErr)
		return
	}
	permId, permErr := strconv.Atoi(permissionIdStr)
	if permErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid permission Id passed", permErr)
		return
	}

	err := rc.roleService.RemovePermission(roleId, permId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while removing perm", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Successfully removed permission to role", nil)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request) {
	roleIdString := chi.URLParam(r, "roleId")
	roleId, roleErr := strconv.Atoi(roleIdString)
	if roleErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid roleId passed", roleErr)
		return
	}

	permissions, err := rc.roleService.GetRolePermissions(roleId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "something went wrong while fetching perm", err)
		return
	}

	if permissions == nil {
		utils.WriteJsonSucessResponse(w, http.StatusNotFound, "no permissions found", nil)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "fetched all permissions", permissions)
}

func (rc *RoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	roleIdStr := chi.URLParam(r, "roleId")

	userId, userIdErr := strconv.Atoi(userIdStr)
	if userIdErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "wrong userId passed", userIdErr)
		return
	}

	roleId, roleIdErr := strconv.Atoi(roleIdStr)
	if roleIdErr != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "wrong role id passed", roleIdErr)
		return
	}

	err := rc.roleService.AssignRoleToUser(userId, roleId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "error occured while assigning role to the user", err)
		return
	}

	utils.WriteJsonSucessResponse(w, http.StatusOK, "Successfully assigned role to the user", nil)

}