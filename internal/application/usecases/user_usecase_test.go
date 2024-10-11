// internal/application/usecases/user_usecase_test.go
package usecases_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/auction-system/internal/application/usecases"
	"github.com/yourusername/auction-system/internal/domain/entities"
	"testing"
)

type MockUserRepository struct {
	users map[int64]*entities.User
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	user.ID = int64(len(m.users) + 1)
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	return m.users[id], nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *entities.User) error {
	m.users[user.ID] = user
	return nil
}

func TestCreateUser(t *testing.T) {
	repo := &MockUserRepository{users: make(map[int64]*entities.User)}
	useCase := usecases.NewUserUseCase(repo)
	user := &entities.User{Name: "John Doe"}
	err := useCase.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), user.ID)
}
