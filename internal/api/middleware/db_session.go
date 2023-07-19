package middleware

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/daos"
)

func DBSessionMiddleware(dao *daos.Dao) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dao", dao)
			return next(c)
		}
	}
}
