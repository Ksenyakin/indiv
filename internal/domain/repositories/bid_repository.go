// internal/domain/repositories/bid_repository.go

package repositories

import (
	"context"
	"indiv/internal/domain/entities"
)

// BidRepository определяет методы для доступа к данным ставок.
type BidRepository interface {
	// Create добавляет новую ставку в репозиторий.
	Create(ctx context.Context, bid *entities.Bid) error

	// GetByAuctionID получает все ставки для конкретного аукциона.
	GetByAuctionID(ctx context.Context, auctionID int64) ([]*entities.Bid, error)

	// GetHighestBid получает самую высокую ставку для конкретного аукциона.
	GetHighestBid(ctx context.Context, auctionID int64) (*entities.Bid, error)

	// RefundLosingBids возвращает средства пользователям, которые не выиграли аукцион.
	RefundLosingBids(ctx context.Context, auctionID, winnerID int64) error
}
