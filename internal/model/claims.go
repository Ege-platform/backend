package model

type Claims struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

type VKClaims struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserID      int    `json:"user_id"`
}

type TGClaims struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
}
