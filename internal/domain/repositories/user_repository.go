// internal/domain/repositories/user_repository.go

package repositories

import (
	"context"
	"indiv/internal/domain/entities"
)

// UserRepository определяет методы для доступа к данным пользователей.
type UserRepository interface {
	// Create добавляет нового пользователя в репозиторий.
	Create(ctx context.Context, user *entities.User) error

	// GetByID получает пользователя по его ID.
	GetByID(ctx context.Context, id int64) (*entities.User, error)

	// Update обновляет информацию существующего пользователя.
	Update(ctx context.Context, user *entities.User) error
}
