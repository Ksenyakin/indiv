// pkg/database/database.go
package database

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"indiv/pkg/config"
)

func NewPostgresConnection(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	return sql.Open("postgres", dsn)
}
