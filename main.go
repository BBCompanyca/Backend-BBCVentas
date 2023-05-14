package main

import (
	"context"
	"fmt"

	"github.com/BBCompanyca/Backend-BBCVentas.git/database"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/service"
	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
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
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "neiferjr14@gmail.com", "Neifer", "Dilanjr15,.")
				if err != nil {
					panic(err)
				}

				u, err := serv.Login(ctx, "neiferjr14@gmail.com", "Dilanjr15,.")

				fmt.Println(u)

				if u.Name != "Neifer" {
					panic("wrong name")
				}

			},
		),
	)

	app.Run()

}
