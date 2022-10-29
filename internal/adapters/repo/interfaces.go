package repo

import (
	"context"

	"github.com/rendau/push/internal/domain/entities"
)

type Repo interface {
	// token
	TokenGet(ctx context.Context, value string) (*entities.TokenSt, error)
	TokenList(ctx context.Context, pars *entities.TokenListParsSt) ([]*entities.TokenSt, int64, error)
	TokenValueExists(ctx context.Context, value string) (bool, error)
	TokenCreate(ctx context.Context, obj *entities.TokenCUSt) (string, error)
	TokenUpdate(ctx context.Context, value string, obj *entities.TokenCUSt) error
	TokenDelete(ctx context.Context, value string) error
}
