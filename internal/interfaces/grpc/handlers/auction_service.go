// internal/interfaces/grpc/handlers/auction_service.go

package handlers

import (
	"context"
	"time"

	"indiv/internal/application/usecases"
	auctionpb "indiv/proto/v1/auction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuctionServiceServer struct {
	useCase *usecases.AuctionUseCase
	auctionpb.UnimplementedAuctionServiceServer
}

func NewAuctionServiceServer(useCase *usecases.AuctionUseCase) *AuctionServiceServer {
	return &AuctionServiceServer{useCase: useCase}
}

// GetAuctionByID получает аукцион по ID
func (s *AuctionServiceServer) GetAuctionByID(ctx context.Context, req *auctionpb.GetAuctionByIDRequest) (*auctionpb.GetAuctionByIDResponse, error) {
	auction, err := s.useCase.GetAuctionByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения аукциона: %v", err)
	}
	if auction == nil {
		return nil, status.Errorf(codes.NotFound, "Аукцион не найден")
	}

	return &auctionpb.GetAuctionByIDResponse{
		Auction: &auctionpb.Auction{
			Id:           auction.ID,
			LotId:        auction.LotID,
			Status:       auction.Status,
			WinnerId:     auction.WinnerID,
			FinalPrice:   auction.FinalPrice,
			AuctionStart: auction.AuctionStart.Format(time.RFC3339),
			AuctionEnd:   auction.AuctionEnd.Format(time.RFC3339),
		},
	}, nil
}

// CloseAuction закрывает аукцион
func (s *AuctionServiceServer) CloseAuction(ctx context.Context, req *auctionpb.CloseAuctionRequest) (*auctionpb.CloseAuctionResponse, error) {
	auction, err := s.useCase.CloseAuction(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка закрытия аукциона: %v", err)
	}

	return &auctionpb.CloseAuctionResponse{
		Auction: &auctionpb.Auction{
			Id:           auction.ID,
			LotId:        auction.LotID,
			Status:       auction.Status,
			WinnerId:     auction.WinnerID,
			FinalPrice:   auction.FinalPrice,
			AuctionStart: auction.AuctionStart.Format(time.RFC3339),
			AuctionEnd:   auction.AuctionEnd.Format(time.RFC3339),
		},
	}, nil
}
