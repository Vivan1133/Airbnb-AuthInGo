package dtos

type CreateUserRequestDto struct {
	Name string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type CreaUserResponseDto struct {
	Message string `json:"message"`
	StatusCode int64 `json:"statusCode"`
}

