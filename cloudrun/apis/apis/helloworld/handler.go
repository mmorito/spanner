package helloworld

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	var uc UseCase
	res, _ := uc.helloworld()
	return c.JSON(http.StatusOK, res)
}
