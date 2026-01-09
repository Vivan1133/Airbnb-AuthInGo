package dtos

type CreateRoleDTO struct {
	Name 			string		`json:"name"`
	Description		string		`json:"description"`
}

type UpdateRoleDTO struct {
	Id				string		`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
}

type PermissionRequestDTO struct {
	Name     		string 		`json:"name"`
	Desc     		string 		`json:"desc"`
	Resource 		string 		`json:"resource"`
	Action   		string 		`json:"action"`
}