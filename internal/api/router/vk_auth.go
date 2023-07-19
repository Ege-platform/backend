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
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error(), Err: err})
		return err
	}
	c.JSON(http.StatusOK, claims)
	return nil
}
