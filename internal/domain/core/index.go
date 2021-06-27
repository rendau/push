package core

import (
	"context"

	"github.com/rendau/push/internal/domain/entities"
	"github.com/rendau/push/internal/errs"
	"github.com/rendau/push/internal/interfaces"
)

const sessionContextKey = "usr_session"

type St struct {
	lg           interfaces.Logger
	db           interfaces.Db
	fcmServerKey string
}

func New(lg interfaces.Logger, db interfaces.Db, fcmServerKey string) *St {
	return &St{
		lg:           lg,
		db:           db,
		fcmServerKey: fcmServerKey,
	}
}

func (c *St) ContextWithSession(ctx context.Context, ses *entities.Session) context.Context {
	return context.WithValue(ctx, sessionContextKey, ses)
}

func (c *St) ContextGetSession(ctx context.Context) *entities.Session {
	contextV := ctx.Value(sessionContextKey)
	if contextV == nil {
		return &entities.Session{}
	}

	switch ses := contextV.(type) {
	case *entities.Session:
		return ses
	default:
		c.lg.Fatal("wrong type of session in context")
		return nil
	}
}

func (c *St) SesRequireAuth(ses *entities.Session) error {
	if ses.Id == 0 {
		return errs.NotAuthorized
	}
	return nil
}
