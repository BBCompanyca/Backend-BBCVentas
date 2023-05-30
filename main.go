package main

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(

		fx.Provide(

			context.Background,
		),
		//signin.SigninModule,
		user.UserModule,
	)
	app.Run()
}
