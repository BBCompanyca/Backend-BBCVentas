package main

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/database"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(

			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from user")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()

}
