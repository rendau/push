package interfaces

import "github.com/rendau/push/internal/domain/entities"

type Db interface {
	CreateToken(st *entities.TokenCreateSt) error
	DeleteToken(token string) error
	GetTokens(platformId int, usrIds []int64) ([]string, error)
	DeleteTokens(tokens []string) error
	DeleteUsr(usrId int64) error
}
