package pg

import (
	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	db.RDBConnectionWithHelpers

	lg logger.Lite
}

func New(db db.RDBConnectionWithHelpers, lg logger.Lite) *St {
	return &St{
		RDBConnectionWithHelpers: db,

		lg: lg,
	}
}
