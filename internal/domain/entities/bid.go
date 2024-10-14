// internal/domain/entities/bid.go

package entities

import (
	"indiv/internal/infrastructure/repositories"
	"time"
)

type Bid struct {
	ID        int64     `json:"id"`
	AuctionID int64     `json:"auction_id"`
	BidderID  int64     `json:"bidder_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

func NewBidUseCase(
	bidRepo repositories.BidRepository,
	auctionRepo repositories.AuctionRepository,
	userRepo repositories.UserRepository,
	lotRepo repositories.LotRepository,
) *BidUseCase {
	return &BidUseCase{
		bidRepo:     bidRepo,
		auctionRepo: auctionRepo,
		userRepo:    userRepo,
		lotRepo:     lotRepo,
	}
}
