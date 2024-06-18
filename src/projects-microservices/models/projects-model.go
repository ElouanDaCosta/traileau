package projects_model

type Project struct {
	Name        string `json:"name" validate:"required" binding:"required"`
	Description string `json:"description" validate:"required" binding:"required"`
	Author      string `validate:"required" binding:"required"`
}
