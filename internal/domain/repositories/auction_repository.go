// internal/domain/repositories/auction_repository.go

package repositories

import (
	"context"
	"indiv/internal/domain/entities"
	"time"
)

// AuctionRepository определяет методы для доступа к данным аукционов.
type AuctionRepository interface {
	// GetByID получает аукцион по его ID.
	GetByID(ctx context.Context, id int64) (*entities.Auction, error)

	// Update обновляет информацию существующего аукциона.
	Update(ctx context.Context, auction *entities.Auction) error

	// GetAuctionsEndingBefore получает аукционы, которые заканчиваются до определенного времени и все еще открыты.
	GetAuctionsEndingBefore(ctx context.Context, endTime time.Time) ([]*entities.Auction, error)
}
