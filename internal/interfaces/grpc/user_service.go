// internal/interfaces/grpc/user_service.go
package grpc

import (
	"context"
	pb "github.com/yourusername/auction-system/api/proto"
	"github.com/yourusername/auction-system/internal/application/usecases"
	"github.com/yourusername/auction-system/internal/domain/entities"
)

type UserServiceServer struct {
	useCase *usecases.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(useCase *usecases.UserUseCase) *UserServiceServer {
	return &UserServiceServer{useCase: useCase}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &entities.User{Name: req.Name}
	if err := s.useCase.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: &pb.User{
		Id:      user.ID,
		Name:    user.Name,
		Balance: user.Balance,
	}}, nil
}

// Реализуйте остальные методы аналогично
