// internal/interfaces/grpc/handlers/bid_service.go

package handlers

import (
	"context"
	"time"

	"indiv/internal/application/usecases"
	"indiv/internal/domain/entities"
	bidpb "indiv/proto/v1/bid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BidServiceServer struct {
	useCase *usecases.BidUseCase
	bidpb.UnimplementedBidServiceServer
}

func NewBidServiceServer(useCase *usecases.BidUseCase) *BidServiceServer {
	return &BidServiceServer{useCase: useCase}
}

// PlaceBid принимает ставку
func (s *BidServiceServer) PlaceBid(ctx context.Context, req *bidpb.PlaceBidRequest) (*bidpb.PlaceBidResponse, error) {
	bid := &entities.Bid{
		AuctionID: req.AuctionId,
		BidderID:  req.BidderId,
		Amount:    req.Amount,
		Timestamp: time.Now(),
	}

	err := s.useCase.PlaceBid(ctx, bid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка размещения ставки: %v", err)
	}

	return &bidpb.PlaceBidResponse{
		Bid: &bidpb.Bid{
			Id:        bid.ID,
			AuctionId: bid.AuctionID,
			BidderId:  bid.BidderID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp.Format(time.RFC3339),
		},
	}, nil
}

// GetBidsByAuction получает ставки по аукциону
func (s *BidServiceServer) GetBidsByAuction(ctx context.Context, req *bidpb.GetBidsByAuctionRequest) (*bidpb.GetBidsByAuctionResponse, error) {
	bids, err := s.useCase.GetBidsByAuction(ctx, req.AuctionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения ставок: %v", err)
	}

	var bidList []*bidpb.Bid
	for _, bid := range bids {
		bidList = append(bidList, &bidpb.Bid{
			Id:        bid.ID,
			AuctionId: bid.AuctionID,
			BidderId:  bid.BidderID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp.Format(time.RFC3339),
		})
	}

	return &bidpb.GetBidsByAuctionResponse{
		Bids: bidList,
	}, nil
}
