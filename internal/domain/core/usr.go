package core

import (
	"context"

	"github.com/rendau/push/internal/domain/entities"
)

type Usr struct {
	r *St
}

func NewUsr(r *St) *Usr {
	return &Usr{r: r}
}

func (c *Usr) TokenDestroy(ctx context.Context, usrId int64) error {
	tokens, _, err := c.r.Token.list(ctx, &entities.TokenListParsSt{
		UsrId: &usrId,
	})
	if err != nil {
		return err
	}

	for _, token := range tokens {
		err = c.r.Token.Delete(ctx, token.Value)
		if err != nil {
			return err
		}
	}

	return nil
}
