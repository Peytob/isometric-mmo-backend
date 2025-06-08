package init

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"isonetric-mmo-backend/internal/model"
)

func SqlDatabase(config *model.SqlConfig) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", buildConnectionString(config))
}

func buildConnectionString(config *model.SqlConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DbName, config.Password, "disable")
}
