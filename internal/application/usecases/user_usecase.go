// internal/application/usecases/user_usecase.go

package usecases

import (
	"context"
	"errors"
	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
	"indiv/internal/infrastructure/adapters"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
	payment  adapters.PaymentAdapter
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
	// Создание платежа
	payment := &entities.Payment{
		UserID: userID,
		Amount: amount,
		Status: "PENDING",
	}

	// Обработка платежа через адаптер
	err := uc.payment.ProcessPayment(ctx, payment)
	if err != nil {
		payment.Status = "FAILED"
		// Логирование ошибки или сохранение информации о неудачном платеже
		return err
	}

	// Обновление баланса пользователя после успешного платежа
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("пользователь не найден")
	}

	user.Balance += amount
	return uc.userRepo.Update(ctx, user)
}
