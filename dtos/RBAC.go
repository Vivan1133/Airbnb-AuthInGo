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