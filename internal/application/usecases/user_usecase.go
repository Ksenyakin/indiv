package usecases

import (
	"context"
	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *entities.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *entities.User) error {
	return uc.userRepo.Update(ctx, user)
}

func (uc *UserUseCase) TopUpBalance(ctx context.Context, userID int64, amount float64) error {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil || user == nil {
		return err
	}
	user.Balance += amount
	return uc.userRepo.Update(ctx, user)
}
