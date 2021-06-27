package pg

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/rendau/push/internal/domain/entities"
)

func (db *St) CreateToken(parse *entities.TokenCreateSt) error {
	var err error
	var dbUsrId int64
	var dbPlatformId int

	err = db.Db.QueryRowx(`
		select usr_id, platform_id 
		from usr_push_token
		where value=$1
	`, parse.Token).Scan(&dbUsrId, &dbPlatformId)
	if err != sql.ErrNoRows {
		if err != nil {
			return err
		}
		if dbUsrId != parse.Id || dbPlatformId != parse.PlatformId {
			_, err = db.Db.Exec(`
				update usr_push_token
				set usr_id=$1, platform_id=$2
				where value=$3
			`, parse.Id, parse.PlatformId, parse.Token)
			if err != nil {
				return err
			}
		}
		return nil
	}

	_, err = db.Db.Exec(`
			insert into usr_push_token(value, platform_id, usr_id) 
			values($1, $2, $3)
	`, parse.Token, parse.PlatformId, parse.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *St) DeleteToken(token string) error {
	_, err := db.Db.Exec(`
		delete from usr_push_token 
		where value=$1
	`, token)
	if err != nil {
		return err
	}

	return nil
}

func (db *St) GetTokens(platformId int, usrIds []int64) ([]string, error) {
	var tokens []string

	err := db.Db.Select(
		&tokens,
		`
			select value
			from usr_push_token
			where usr_id = any($1) and platform_id=$2
		`,
		pq.Array(usrIds),
		platformId,
	)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (db *St) DeleteTokens(tokens []string) error { // delete tokens
	_, err := db.Db.Exec(`
		delete from usr_push_token 
		where value = any($1)
	`, pq.Array(tokens))
	if err != nil {
		return err
	}

	return nil
}

func (db *St) DeleteUsr(usrId int64) error {
	_, err := db.Db.Exec(`
		delete from usr_push_token 
		where usr_id=$1
	`, usrId)
	if err != nil {
		return err
	}

	return nil
}
