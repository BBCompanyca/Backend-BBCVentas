package database

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	connetionString := "root:" + "" + "@tcp(127.0.0.1:3306)/test_bbcv?parseTime=true"

	/*connetionString := fmt.Sprintf(
		"%s:+%s+@tcp(%s:%d)/%s",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)*/

	return sqlx.ConnectContext(ctx, "mysql", connetionString)

}
