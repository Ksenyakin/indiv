// internal/interfaces/migration/migrate.go

package migration

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"indiv/pkg/config"
)

func Migrate(db *sql.DB) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("ошибка загрузки конфигурации: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("ошибка создания драйвера миграции: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		cfg.Database.DBName,
		driver,
	)
	if err != nil {
		return fmt.Errorf("ошибка создания экземпляра миграции: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("ошибка применения миграции: %v", err)
	}

	return nil
}
