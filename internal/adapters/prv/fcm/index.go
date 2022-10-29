package fcm

import (
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg logger.Lite
}

func New(lg logger.Lite) *St {
	return &St{
		lg: lg,
	}
}
