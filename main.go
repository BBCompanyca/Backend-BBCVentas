package main

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/database"
	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),
		fx.Invoke(),
	)

	app.Run()

}
