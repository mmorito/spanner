package middleware

import (
	"github.com/labstack/echo/v4"

	"github.com/mmorito/spanner/utilities/logger"
)

// SetupLogger middleware
func SetupLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.SetupLogger(c.Request())
		return next(c)
	}
}
