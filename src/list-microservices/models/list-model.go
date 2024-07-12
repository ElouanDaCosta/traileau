package list_model

type List struct {
	Name     string `json:"name" validate:"required" binding:"required"`
	Position int    `json:"position" validate:"required" binding:"required"`
}
