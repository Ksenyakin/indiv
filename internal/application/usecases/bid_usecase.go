// internal/application/usecases/bid_usecase.go

package usecases

import (
	"context"
	"errors"

	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
)

type BidUseCase struct {
	bidRepo     repositories.BidRepository
	auctionRepo repositories.AuctionRepository
	userRepo    repositories.UserRepository
}

func NewBidUseCase(bidRepo repositories.BidRepository, auctionRepo repositories.AuctionRepository, userRepo repositories.UserRepository) *BidUseCase {
	return &BidUseCase{
		bidRepo:     bidRepo,
		auctionRepo: auctionRepo,
		userRepo:    userRepo,
	}
}

func (uc *BidUseCase) PlaceBid(ctx context.Context, bid *entities.Bid) error {
	// Проверка аукциона
	auction, err := uc.auctionRepo.GetByID(ctx, bid.AuctionID)
	if err != nil {
		return err
	}
	if auction == nil || auction.Status != "OPEN" {
		return errors.New("Аукцион не открыт или не существует")
	}

	// Получение лота
	lot, err := uc.lotRepo.GetByID(ctx, auction.LotID)
	if err != nil {
		return err
	}
	if lot == nil {
		return errors.New("Лот не найден")
	}

	// Проверка баланса пользователя
	user, err := uc.userRepo.GetByID(ctx, bid.BidderID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("Пользователь не найден")
	}
	if user.Balance < bid.Amount {
		return errors.New("Недостаточно средств")
	}

	// Проверка минимального шага
	highestBid, err := uc.bidRepo.GetHighestBid(ctx, bid.AuctionID)
	minBid := lot.StartingPrice
	if highestBid != nil {
		minBid = highestBid.Amount + lot.MinBidIncrement
	}
	if bid.Amount < minBid {
		return errors.New("Сумма ставки недостаточна")
	}

	// Размещение ставки
	return uc.bidRepo.Create(ctx, bid)
}

func (uc *BidUseCase) GetBidsByAuction(ctx context.Context, auctionID int64) ([]*entities.Bid, error) {
	return uc.bidRepo.GetByAuctionID(ctx, auctionID)
}
