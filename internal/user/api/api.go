package api

import (
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv          service.Service
	dataValidator *validator.Validate
}

func New(serv service.Service) *API {
	return &API{
		serv:          serv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	return e.Start(address)
}
