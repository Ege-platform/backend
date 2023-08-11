package model

type AuthResponse struct {
	Model *User  `json:"model"`
	Token string `json:"token"`
}
