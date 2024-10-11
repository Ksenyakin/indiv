// internal/interfaces/grpc/server.go
package grpc

import (
	"net"

	auctionpb "indiv/api/proto/auction"
	bidpb "indiv/api/proto/bid"
	lotpb "indiv/api/proto/lot"
	userpb "indiv/api/proto/user"
	"indiv/internal/application/usecases"
	"indiv/internal/interfaces/grpc/handlers"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(port string, userUC *usecases.UserUseCase,
	lotUC *usecases.LotUseCase, bidUC *usecases.BidUseCase,
	auctionUC *usecases.AuctionUseCase, logger *zap.SugaredLogger) error {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server,
		handlers.NewUserServiceServer(userUC))
	lotpb.RegisterLotServiceServer(server,
		handlers.NewLotServiceServer(lotUC))
	bidpb.RegisterBidServiceServer(server,
		handlers.NewBidServiceServer(bidUC))
	auctionpb.RegisterAuctionServiceServer(server,
		handlers.NewAuctionServiceServer(auctionUC))

	logger.Infof("gRPC сервер запущен на %s", port)
	return server.Serve(lis)
}
