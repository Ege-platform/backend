package model

type LoginCredentials struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}
