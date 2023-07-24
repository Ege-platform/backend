package router

import (
	"ege_platform/internal/api/controller"
	"ege_platform/internal/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (r *Router) AuthWithVK(c echo.Context) error {
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(r.Cfg.AuthURI, r.Cfg.ClientID, r.Cfg.RedirectURI))
	return nil
}

func (r *Router) InternalAuth(c echo.Context) error {
	claims, err := controller.GetVKClaims(c, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: err.Error(), Err: err})
	}

	loginCredentials, err := controller.GenerateLoginCredentials(c, claims, r.Cfg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error(), Err: err})
	}

	return c.JSON(http.StatusOK, loginCredentials)
}
