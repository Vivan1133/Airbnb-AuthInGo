package controllers

import (
	"AuthInGo/dtos"
	"AuthInGo/services"
	"AuthInGo/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PermissionController struct {
	permissionService services.IPermissionService
}

func NewPermissionController(permissionService services.IPermissionService) (*PermissionController) {
	return &PermissionController {
		permissionService : permissionService,
	}
}

func (pc *PermissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var req dtos.PermissionRequestDTO

	if err := utils.ReadJsonRequest(r, &req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	permission, err := pc.permissionService.CreatePermission(
		req.Name,
		req.Desc,
		req.Resource,
		req.Action,
	)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to create permission", err)
		return
	}

	utils.WriteJsonSucessResponse(
		w,
		http.StatusCreated,
		"permission created successfully",
		permission,
	)
}

func (pc *PermissionController) GetPermissionById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid permission id", err)
		return
	}

	permission, err := pc.permissionService.GetPermissionById(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "permission not found", err)
		return
	}

	utils.WriteJsonSucessResponse(
		w,
		http.StatusOK,
		"permission fetched successfully",
		permission,
	)
}

func (pc *PermissionController) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := pc.permissionService.GetAllPermissions()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to fetch permissions", err)
		return
	}

	utils.WriteJsonSucessResponse(
		w,
		http.StatusOK,
		"permissions fetched successfully",
		permissions,
	)
}

func (pc *PermissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid permission id", err)
		return
	}

	var req dtos.PermissionRequestDTO
	if err := utils.ReadJsonRequest(r, &req); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	permission, err := pc.permissionService.UpdatePermission(
		id,
		req.Name,
		req.Desc,
		req.Resource,
		req.Action,
	)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to update permission", err)
		return
	}

	utils.WriteJsonSucessResponse(
		w,
		http.StatusOK,
		"permission updated successfully",
		permission,
	)
}

func (pc *PermissionController) DeletePermissionById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid permission id", err)
		return
	}

	if err := pc.permissionService.DeletePermissionById(id); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to delete permission", err)
		return
	}

	utils.WriteJsonSucessResponse(
		w,
		http.StatusOK,
		"permission deleted successfully",
		nil,
	)
}
