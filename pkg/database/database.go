// pkg/database/database.go
package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yourusername/auction-system/pkg/config"
)

func NewPostgresConnection(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	return sql.Open("postgres", dsn)
}
