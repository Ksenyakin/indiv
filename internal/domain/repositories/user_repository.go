// internal/domain/repositories/user_repository.go
package repositories

import (
	"context"
	"indiv/internal/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByID(ctx context.Context, id int64) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
}
