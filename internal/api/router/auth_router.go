package router

import (
	"ege_platform/internal/api/controller"
	"ege_platform/internal/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (r *Router) AuthWithVK(c echo.Context) error {
	return c.JSON(http.StatusOK, fmt.Sprintf(r.Cfg.VKAuthURI, r.Cfg.VKClientID, r.Cfg.VKRedirectURI))
}

func (r *Router) InternalVKAuth(c echo.Context) error {
	accessToken, err := controller.AuthWithVK(c, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	return c.Redirect(http.StatusOK, accessToken)
}

func (r *Router) AuthWithTG(c echo.Context) error {
	accessToken, err := controller.AuthWithTG(c, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	return c.Redirect(http.StatusOK, accessToken)
}

func (r *Router) GetAccessToken(c echo.Context) error {
	accessToken, err := controller.GetAccessToken(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrorResponse{Err: err.Error()})
	}

	return c.JSON(http.StatusOK, accessToken)
}
