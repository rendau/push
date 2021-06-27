package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/rendau/push/internal/domain/core"
	"github.com/rendau/push/internal/interfaces"
)

type St struct {
	lg         interfaces.Logger
	cr         *core.St
	usrAuthUrl string

	server *http.Server
	lChan  chan error
}

func New(lg interfaces.Logger, cr *core.St, usrAuthUrl string, listen string) *St {
	api := &St{
		lg:         lg,
		cr:         cr,
		usrAuthUrl: usrAuthUrl,
		lChan:      make(chan error, 1),
	}

	api.server = &http.Server{
		Addr:              listen,
		Handler:           api.router(),
		ReadTimeout:       2 * time.Minute,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return api
}

func (a *St) Start() {
	go func() {
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.lg.Errorw("Http server closed", err)
			a.lChan <- err
		}
	}()
}

func (a *St) Wait() <-chan error {
	return a.lChan
}

func (a *St) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
