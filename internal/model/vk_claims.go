package model

type VKClaims struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserID      int    `json:"user_id"`
}
