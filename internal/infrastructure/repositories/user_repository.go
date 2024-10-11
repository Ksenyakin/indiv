// internal/infrastructure/repositories/user_repository.go
package repositories

import (
	"context"
	"database/sql"
	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	query := "INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRowContext(ctx, query, user.Name, user.Balance).Scan(&user.ID)
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	query := "SELECT id, name, balance FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	user := &entities.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Balance)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	query := "UPDATE users SET name = $1, balance = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Balance, user.ID)
	return err
}
