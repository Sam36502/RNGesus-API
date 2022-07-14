package api

import (
	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	registerEndpoints(e)

	return e
}

func registerEndpoints(e *echo.Echo) {

	e.GET("/v1/rand/float", getRandomFloat)
	e.GET("/v1/rand/int", getRandomInt)

}
