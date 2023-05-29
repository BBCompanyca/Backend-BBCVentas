package api

import "github.com/labstack/echo/v4"

func (a *API) Sing_In(e *echo.Echo) {

	e.POST("/sign_in", a.Login)

}
