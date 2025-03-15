package requests

// LoginRequest struct
type LoginRequest struct {
	// add custom validation messages
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
