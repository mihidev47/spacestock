package repository

import (
	"github.com/jmoiron/sqlx"

	"../model"
)

type appRepository struct {
	db *sqlx.DB
}

type AppRepository interface {
	FindConfig(appId string, lastUpdated int64) (cfg []model.AppConfig, err error)
}

func (t *appRepository) FindConfig(appId string, lastUpdated int64) (cfg []model.AppConfig, err error) {
	t.db.Select(&cfg, "SELECT * FROM app_config WHERE updated_at > FROM_UNIXTIME(?) AND app_id = ?",
		lastUpdated, appId)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return cfg, err
}
