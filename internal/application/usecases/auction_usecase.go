// internal/application/usecases/auction_usecase.go

package usecases

import (
	"context"
	"errors"
	"time"

	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
)

type AuctionUseCase struct {
	auctionRepo repositories.AuctionRepository
	bidRepo     repositories.BidRepository
	userRepo    repositories.UserRepository
}

func NewAuctionUseCase(auctionRepo repositories.AuctionRepository, bidRepo repositories.BidRepository, userRepo repositories.UserRepository) *AuctionUseCase {
	return &AuctionUseCase{
		auctionRepo: auctionRepo,
		bidRepo:     bidRepo,
		userRepo:    userRepo,
	}
}

func (uc *AuctionUseCase) GetAuctionByID(ctx context.Context, id int64) (*entities.Auction, error) {
	return uc.auctionRepo.GetByID(ctx, id)
}

func (uc *AuctionUseCase) CloseAuction(ctx context.Context, id int64) (*entities.Auction, error) {
	auction, err := uc.auctionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if auction == nil {
		return nil, errors.New("Аукцион не найден")
	}
	if auction.Status != "OPEN" {
		return nil, errors.New("Аукцион уже закрыт")
	}

	highestBid, err := uc.bidRepo.GetHighestBid(ctx, id)
	if err != nil {
		return nil, err
	}

	if highestBid != nil {
		auction.WinnerID = &highestBid.BidderID
		auction.FinalPrice = &highestBid.Amount

		// Списание средств у победителя
		winner, err := uc.userRepo.GetByID(ctx, highestBid.BidderID)
		if err != nil {
			return nil, err
		}
		winner.Balance -= highestBid.Amount
		err = uc.userRepo.Update(ctx, winner)
		if err != nil {
			return nil, err
		}

		// Возврат средств проигравшим
		err = uc.bidRepo.RefundLosingBids(ctx, id, highestBid.BidderID)
		if err != nil {
			return nil, err
		}
	}

	auction.Status = "CLOSED"
	err = uc.auctionRepo.Update(ctx, auction)
	if err != nil {
		return nil, err
	}

	return auction, nil
}

func (uc *AuctionUseCase) GetAuctionsToClose(ctx context.Context) ([]*entities.Auction, error) {
	return uc.auctionRepo.GetAuctionsEndingBefore(ctx, time.Now())
}
