package controller

import (
	"ege_platform/internal/config"
	"ege_platform/internal/logging"
	"ege_platform/internal/model"
	service "ege_platform/internal/service/jwt"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

func AuthWithVK(c echo.Context, cfg *config.Config) (string, error) {
	code := c.QueryParams().Get("code")

	response, err := http.Get(fmt.Sprintf(cfg.VKTokenURI, cfg.VKClientID, cfg.VKClientSecret, cfg.VKRedirectURI, code))
	if err != nil {
		logging.Log.Errorf("Can't get token: %v", err)
		return "", err
	}

	decoder := json.NewDecoder(response.Body)

	vkClaims := &model.VKClaims{}
	if err := decoder.Decode(vkClaims); err != nil {
		logging.Log.Errorf("Can't parse json: %v", err)
		return "", err
	}

	if vkClaims.AccessToken == "" {
		logging.Log.Debug("Empty claims")
		return "", errors.New("unauthorized")
	}

	vk := api.NewVK(vkClaims.AccessToken)

	userVK, err := vk.UsersGet(api.Params{
		"user_ids": vkClaims.UserID,
	})
	if err != nil {
		logging.Log.Errorf("Can't get user name from VK: %v", err)
		return "", err
	}
	claims := &model.Claims{
		Username:    fmt.Sprintf("%d", vkClaims.UserID) + "VK",
		AccessToken: vkClaims.AccessToken,
		Name:        userVK[0].FirstName + " " + userVK[0].LastName,
	}

	return service.AuthenticateUser(c, claims, cfg.JwtSecret), nil
}

func AuthWithTG(c echo.Context, cfg *config.Config) (string, error) {
	decoder := json.NewDecoder(c.Request().Body)

	tgClaims := &model.TGClaims{}
	if err := decoder.Decode(tgClaims); err != nil {
		logging.Log.Errorf("Can't read body: %v", err)
		return "", err
	}

	claims := &model.Claims{
		Username:    fmt.Sprintf("%d", tgClaims.ID) + "TG",
		AccessToken: "123",
		Name:        tgClaims.FirstName + " " + tgClaims.LastName,
	}

	return service.AuthenticateUser(c, claims, cfg.JwtSecret), nil
}

func GetAccessToken(c echo.Context) (string, error) {
	dao := c.Get("dao").(*daos.Dao)
	username := c.QueryParams().Get("username")

	userPB, err := dao.FindAuthRecordByUsername("users", username)
	if err != nil {
		logging.Log.Errorf("No user with %s username: %v", username, err)
		return "", err
	}

	return userPB.TokenKey(), nil
}
