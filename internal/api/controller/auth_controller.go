package controller

import (
	"ege_platform/internal/config"
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/labstack/echo/v5"
)

func GetClaimsFromVK(c echo.Context, cfg *config.Config) (*model.Claims, error) {
	code := c.QueryParams().Get("code")

	response, err := http.Get(fmt.Sprintf(cfg.VKTokenURI, cfg.VKClientID, cfg.VKClientSecret, cfg.VKRedirectURI, code))
	if err != nil {
		logging.Log.Errorf("Can't get token: %v", err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logging.Log.Errorf("Error while reading response body: %v", err)
		return nil, err
	}

	vkClaims := &model.VKClaims{}
	err = json.Unmarshal(body, &vkClaims)
	if err != nil {
		logging.Log.Errorf("Can't parse json: %v", err)
		return nil, err
	}

	if vkClaims.AccessToken == "" {
		logging.Log.Debug("Empty claims")
		return nil, errors.New("unauthorized")
	}

	vk := api.NewVK(vkClaims.AccessToken)

	userVK, err := vk.UsersGet(api.Params{
		"user_ids": vkClaims.UserID,
	})
	if err != nil {
		logging.Log.Errorf("Can't get user name from VK: %v", err)
		return nil, err
	}

	logging.Log.Debugf("Access token from %d: %s", vkClaims.UserID, vkClaims.AccessToken)

	return &model.Claims{
		ID:          vkClaims.UserID,
		AccessToken: vkClaims.AccessToken,
		Name:        userVK[0].FirstName + " " + userVK[0].LastName,
	}, nil
}

func GetClaimsFromTG(c echo.Context, cfg *config.Config) (*model.Claims, error) {
	decoder := json.NewDecoder(c.Request().Body)

	tgClaims := &model.TGClaims{}
	if err := decoder.Decode(tgClaims); err != nil {
		logging.Log.Errorf("Can't read body: %v", err)
		return nil, err
	}

	return &model.Claims{
		ID:          tgClaims.ID,
		AccessToken: "123",
		Name:        tgClaims.FirstName + " " + tgClaims.LastName,
	}, nil
}
