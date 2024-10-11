package grpc

import (
	"context"
	"indiv/internal/application/usecases"
)

type AuctionServer struct {
	AuctionUseCase usecases.AuctionUseCase
}

func (s *AuctionServer) PlaceBid(ctx context.Context, req *pb.PlaceBidRequest) (*pb.PlaceBidResponse, error) {
	// Обработка gRPC запроса
}
