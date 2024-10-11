package repositories

import (
	"database/sql"
	"indiv/internal/domain/entities"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (p *PostgresUserRepository) Create(user *entities.User) error {
	// Реализация создания пользователя в базе данных
}
