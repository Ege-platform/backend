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

func (r *Router) InternalAuth(c echo.Context) error {
	claims, err := controller.GetClaimsFromVK(c, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	err = controller.CreateUser(c, claims)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, r.Cfg.BaseURL)
}

func (r *Router) AuthWithTG(c echo.Context) error {
	claims, err := controller.GetClaimsFromTG(c, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	err = controller.CreateUser(c, claims)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Err: err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, r.Cfg.BaseURL)
}
