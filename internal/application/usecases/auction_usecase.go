// auction_usecase.go
package usecases

import "indiv/internal/domain/repositories"

type AuctionUseCase struct {
	AuctionRepo repositories.AuctionRepository
	BidRepo     repositories.BidRepository
	// Дополнительные зависимости
}

func (a *AuctionUseCase) PlaceBid(bid *entities.Bid) error {
	// Логика приема ставки
}
