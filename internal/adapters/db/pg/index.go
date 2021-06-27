package pg

import (
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // driver
	"github.com/jmoiron/sqlx"
	"github.com/rendau/push/internal/interfaces"
)

type St struct {
	lg interfaces.Logger
	Db *sqlx.DB
}

func New(lg interfaces.Logger, dsn string) (*St, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(10 * time.Minute)

	return &St{
		lg: lg,
		Db: db,
	}, nil
}
