package init

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"isonetric-mmo-backend/internal/app"
	"isonetric-mmo-backend/internal/model"
)

func Application(config *model.Config, db *sqlx.DB) (*app.Application, error) {
	var err error

	stores, err := store(config, db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize stores: %w", err)
	}

	services, err := service(config, stores)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize services: %w", err)
	}

	return app.NewApplication(services, stores), nil
}

func store(config *model.Config, db *sqlx.DB) (*app.Store, error) {
	return &app.Store{}, nil
}

func service(config *model.Config, store *app.Store) (*app.Service, error) {
	return &app.Service{}, nil
}
