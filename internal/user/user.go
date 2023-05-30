package user

import (
	"context"
	"fmt"

	"github.com/BBCompanyca/Backend-BBCVentas.git/database"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/api"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var UserModule = fx.Module("user",

	fx.Provide(

		settings.New,
		database.New,

		repository.New,
		service.New,
		api.New,

		echo.New,
	),

	fx.Invoke(SetLifeCycleUser),
)

func SetLifeCycleUser(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
