package api

import (
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{serv}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}
