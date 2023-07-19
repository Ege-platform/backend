package middleware

import (
	"ege_platform/internal/logging"
	"fmt"

	"github.com/labstack/echo/v5"
	mw "github.com/labstack/echo/v5/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return mw.RequestLoggerWithConfig(mw.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, values mw.RequestLoggerValues) error {
			logging.Log.Debug(fmt.Sprintf("%s - %s - %d - %d", values.Method, values.URI, values.Latency, values.Status))
			return nil
		},
	})
}
