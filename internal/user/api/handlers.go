package api

import (
	"errors"
	"net/http"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/api/dtos"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.SaveUser(ctx, params.Name, params.Username, params.Password, params.Permissions, 1, "0000-00-00", "0000-00-00", params.Registered_By)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, err)
		}

		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))

	}

	return c.JSON(http.StatusCreated, nil)

}

func (a *API) GetAllUser(c echo.Context) error {

	ctx := c.Request().Context()

	u, err := a.serv.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusOK, u)
}
