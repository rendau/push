package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/push/internal/domain/entities"
)

func (d *St) TokenGet(ctx context.Context, value string) (*entities.TokenSt, error) {
	result := &entities.TokenSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"token"},
		Conds:  []string{"value = ${value}"},
		Args:   map[string]any{"value": value},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) TokenList(ctx context.Context, pars *entities.TokenListParsSt) ([]*entities.TokenSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.Values != nil {
		conds = append(conds, `t.value in (select * from unnest(${values} :: text[]))`)
		args["values"] = *pars.Values
	}
	if pars.UsrId != nil {
		conds = append(conds, `t.usr_id = ${usr_id}`)
		args["usr_id"] = *pars.UsrId
	}
	if pars.UsrIds != nil {
		conds = append(conds, `t.usr_id in (select * from unnest(${usr_ids} :: bigint[]))`)
		args["usr_ids"] = *pars.UsrIds
	}
	if pars.PlatformId != nil {
		conds = append(conds, `t.platform_id = ${platform_id}`)
		args["platform_id"] = *pars.PlatformId
	}

	result := make([]*entities.TokenSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`token t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.value",
		},
	})

	return result, tCount, err
}

func (d *St) TokenValueExists(ctx context.Context, value string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from token
        where value = $1
    `, value).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) TokenCreate(ctx context.Context, obj *entities.TokenCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "token",
		Obj:    obj,
		RetCol: "value",
		RetV:   &result,
	})

	return result, err
}

func (d *St) TokenUpdate(ctx context.Context, value string, obj *entities.TokenCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "token",
		Obj:   obj,
		Conds: []string{"value = ${cond_id}"},
		Args:  map[string]any{"cond_id": value},
	})
}

func (d *St) TokenDelete(ctx context.Context, value string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "token",
		Conds: []string{"value = ${cond_id}"},
		Args:  map[string]any{"cond_id": value},
	})
}
