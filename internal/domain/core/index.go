package core

import (
	"github.com/rendau/dop/adapters/logger"
	"github.com/rendau/push/internal/adapters/prv"
	"github.com/rendau/push/internal/adapters/repo"
)

type St struct {
	lg   logger.Lite
	repo repo.Repo
	prv  prv.Prv

	Session *Session
	Token   *Token
	Usr     *Usr
}

func New(lg logger.Lite, repo repo.Repo, prv prv.Prv) *St {
	c := &St{
		lg:   lg,
		repo: repo,
		prv:  prv,
	}

	c.Session = NewSession(c)
	c.Token = NewToken(c)
	c.Usr = NewUsr(c)

	return c
}
