// internal/interfaces/grpc/server.go
package grpc

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"indiv/internal/application/usecases"
	"net"
)

func RunServer(port string, userUseCase *usecases.UserUseCase, logger *zap.SugaredLogger) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, NewUserServiceServer(userUseCase))
	logger.Infof("gRPC сервер запущен на %s", port)
	return server.Serve(lis)
}
