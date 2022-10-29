package core

import (
	"context"

	"github.com/rendau/push/internal/cns"
	"github.com/rendau/push/internal/domain/entities"
	"github.com/rendau/push/internal/errs"
)

type Token struct {
	r *St
}

func NewToken(r *St) *Token {
	return &Token{r: r}
}

func (c *Token) ValidateCU(ctx context.Context, obj *entities.TokenCUSt, value string) error {
	forCreate := value == ""

	// Value
	if forCreate && obj.Value == nil {
		return errs.TokenValueRequired
	}
	if obj.Value != nil {
		if *obj.Value == "" || len(*obj.Value) > 300 {
			return errs.BadTokenValue
		}
	}

	// UsrId
	if forCreate && obj.UsrId == nil {
		return errs.UsrIdRequired
	}

	// PlatformId
	if forCreate && obj.PlatformId == nil {
		return errs.PlatformRequired
	}
	if obj.PlatformId != nil {
		if !cns.PlatformIsValid(*obj.PlatformId) {
			return errs.BadPlatform
		}
	}

	return nil
}

func (c *Token) list(ctx context.Context, pars *entities.TokenListParsSt) ([]*entities.TokenSt, int64, error) {
	items, tCount, err := c.r.repo.TokenList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Token) valueExists(ctx context.Context, value string) (bool, error) {
	return c.r.repo.TokenValueExists(ctx, value)
}

func (c *Token) Create(ctx context.Context, obj *entities.TokenCUSt) (string, error) {
	var err error

	ses := c.r.Session.GetFromContext(ctx)
	if ses.Id == 0 {
	}

	if err = c.r.Session.RequireAuth(ses); err != nil {
		return "", err
	}

	obj.UsrId = &ses.Id

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// delete if already exists
	err = c.Delete(ctx, *obj.Value)
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.TokenCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Token) Delete(ctx context.Context, value string) error {
	return c.r.repo.TokenDelete(ctx, value)
}
