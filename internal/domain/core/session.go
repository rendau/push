package core

import (
	"context"
	"strconv"

	"github.com/rendau/dop/adapters/jwt"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/push/internal/domain/entities"
)

const sessionContextKey = "user_session"

type Session struct {
	r *St
}

func NewSession(r *St) *Session {
	return &Session{r: r}
}

func (c *Session) GetFromToken(token string) *entities.Session {
	var session entities.Session

	if jwt.ParsePayload(token, &session) != nil {
		session = entities.Session{}
	}

	session.Id, _ = strconv.ParseInt(session.Sub, 10, 64)

	return &session
}

func (c *Session) SetToContext(ctx context.Context, ses *entities.Session) context.Context {
	return context.WithValue(ctx, sessionContextKey, ses)
}

func (c *Session) SetToContextByToken(ctx context.Context, token string) context.Context {
	return c.SetToContext(ctx, c.GetFromToken(token))
}

func (c *Session) GetFromContext(ctx context.Context) *entities.Session {
	contextV := ctx.Value(sessionContextKey)
	if contextV == nil {
		return &entities.Session{}
	}

	switch ses := contextV.(type) {
	case *entities.Session:
		return ses
	default:
		c.r.lg.Errorw("wrong type of session in context", nil)
		return &entities.Session{}
	}
}

func (c *Session) RequireAuth(ses *entities.Session) error {
	if ses.Id <= 0 {
		return dopErrs.NotAuthorized
	}

	return nil
}
