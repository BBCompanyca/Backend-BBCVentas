package signin

import (
	"context"
	"fmt"

	"github.com/BBCompanyca/Backend-BBCVentas.git/database"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/api"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/repository"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/service"
	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var SigninModule = fx.Module("signin",

	fx.Provide(

		settings.New,
		database.New,

		repository.NewSign_in,
		service.NewSign_in,
		api.NewSign_in,

		echo.New,
	),
	fx.Invoke(SetLifeCycleSignIn),
)

func SetLifeCycleSignIn(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
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
