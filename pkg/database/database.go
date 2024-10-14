// pkg/database/database.go
package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"indiv/pkg/config"
	"time"
)

func NewPostgresConnection(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения с БД: %v", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %v", err)
	}

	return db, nil
}
