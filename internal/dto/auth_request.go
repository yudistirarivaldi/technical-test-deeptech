package dto

type RegisterRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfileRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"required,oneof=L P"`
}
