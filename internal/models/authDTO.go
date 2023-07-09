package models

type AuthDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type PasswordDTO struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
