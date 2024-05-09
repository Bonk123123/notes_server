package dto

type RegistrationDto struct {
	Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    PasswordRepeat string `json:"password_repeat" binding:"required"`
}