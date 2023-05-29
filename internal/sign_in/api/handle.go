package api

import (
	"net/http"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/api/dto"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) Login(c echo.Context) error {

	ctx := c.Request().Context()
	params := dto.Sing_In{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "request invalid"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.Login(ctx, params.Username, params.Password)
	if err != nil {
		if err == service.InvalidCredential {
			return c.JSON(http.StatusNotFound, responseMessage{Message: "invalid credential"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "unxpected error"})

	}

	return c.JSON(http.StatusOK, u)
}
