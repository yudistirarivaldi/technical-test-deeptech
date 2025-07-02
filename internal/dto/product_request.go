package dto

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Stock       string `json:"stock" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Stock       string `json:"stock" validate:"required"`
}
