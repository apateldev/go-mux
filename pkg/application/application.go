// pkg/application/application.go

package application

import (
	"nvxapi2/go.mux/pkg/config"
	"nvxapi2/go.mux/pkg/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}