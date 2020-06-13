package helloworld

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmorito/spanner/utilities/logger"
)

func Get(c echo.Context) error {
	logger.SetupLogger(c.Request())
	var uc UseCase
	res, _ := uc.helloworld()
	return c.JSON(http.StatusOK, res)
}
