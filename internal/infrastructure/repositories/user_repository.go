// internal/infrastructure/repositories/user_repository.go

package repositories

import (
	"context"
	"database/sql"
	"errors"

	"indiv/internal/domain/entities"
	domain "indiv/internal/domain/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
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
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	query := "UPDATE users SET name = $1, balance = $2 WHERE id = $3"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Balance, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Пользователь не найден")
	}
	return nil
}
