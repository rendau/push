package entities

import (
	"time"

	"github.com/rendau/dop/dopTypes"
)

type TokenSt struct {
	Value      string    `json:"value" db:"value"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UsrId      int64     `json:"usr_id" db:"usr_id"`
	PlatformId int       `json:"platform_id" db:"platform_id"`
}

type TokenListParsSt struct {
	dopTypes.ListParams

	Values     *[]int64 `json:"values" form:"values"`
	UsrId      *int64   `json:"usr_id" form:"usr_id"`
	PlatformId *int     `json:"platform_id" form:"platform_id"`
}

type TokenCUSt struct {
	Value      *string `json:"value" db:"value"`
	UsrId      *int64  `json:"usr_id" db:"usr_id"`
	PlatformId *int    `json:"platform_id" db:"platform_id"`
}
