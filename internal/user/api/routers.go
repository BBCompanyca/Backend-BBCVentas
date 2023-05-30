package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/user")

	users.POST("/register", a.RegisterUser)
	users.GET("/all", a.GetAllUser)
	users.GET("/username", a.GetUserByUsername)
}
