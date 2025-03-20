package repository

import (
	"database/sql"
	"errors"
	"streaming-api/internal/model"

	"github.com/labstack/gommon/log"
)

var ErrQuery = errors.New("error find stream key")

type KeysRepository interface {
	FindStreamKey(name, keys string) (*model.Keys, error)
}

type keysRepository struct {
	*sql.DB
}

func NewKeysRepository(db *sql.DB) KeysRepository {
	return &keysRepository{
		db,
	}
}

func (kr *keysRepository) FindStreamKey(name, key string) (*model.Keys, error) {
	keys := &model.Keys{}
	row := kr.QueryRow(`SELECT * FROM "Lives" WHERE "name"=$1 AND "stream_key"=$2`, name, key)

	err := row.Scan(&keys.Name, &keys.Key)

	if err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Keys{}, nil
		}
		return &model.Keys{}, ErrQuery
	}

	return keys, nil
}
