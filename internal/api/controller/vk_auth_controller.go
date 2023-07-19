package controller

import (
	"ege_platform/internal/config"
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetVKClaims(c echo.Context, cfg *config.Config) (*model.VKClaims, error) {
	code := c.QueryParams().Get("code")

	response, err := http.Get(fmt.Sprintf(cfg.TokenURI, cfg.ClientID, cfg.ClientSecret, cfg.RedirectURI, code))
	if err != nil {
		logging.Log.Errorf("Can't get token: %v", err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logging.Log.Errorf("Error while reading response body: %v", err)
		return nil, err
	}

	claims := &model.VKClaims{}
	err = json.Unmarshal(body, &claims)
	if err != nil {
		logging.Log.Errorf("Can't parse json: %v", err)
		return nil, err
	}

	return claims, nil
}
