package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel string
	AppMode  string

	PBAdminEmail    string
	PBAdminPassword string
	PBDebug         bool

	BaseURL string

	VKClientID     string
	VKClientSecret string
	VKRedirectURI  string
	VKAuthURI      string
	VKTokenURI     string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		PBDebug:         os.Getenv("PB_DEBUG") == "true",
		LogLevel:        os.Getenv("LOG_LEVEL"),
		PBAdminEmail:    os.Getenv("PB_ADMIN_EMAIL"),
		PBAdminPassword: os.Getenv("PB_ADMIN_PASSWORD"),
		VKClientID:      os.Getenv("VK_CLIENT_ID"),
		VKClientSecret:  os.Getenv("VK_CLIENT_SECRET"),
		VKRedirectURI:   os.Getenv("VK_REDIRECT_URI"),
		BaseURL:         os.Getenv("BASE_URL"),
		VKAuthURI:       os.Getenv("VK_AUTH_URI"),
		VKTokenURI:      os.Getenv("VK_TOKEN_URI"),
		AppMode:         os.Getenv("APP_MODE"),
	}, nil
}
