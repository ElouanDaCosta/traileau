package users_model

type User struct {
	Username string `json:"username" validate:"required" binding:"required"`
	Email    string `json:"email" validate:"required" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}
