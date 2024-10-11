// internal/infrastructure/repositories/auction_repository.go

package repositories

import (
	"context"
	"database/sql"
	"time"

	"indiv/internal/domain/entities"
	domain "indiv/internal/domain/repositories"
)

type AuctionRepository struct {
	db *sql.DB
}

func NewAuctionRepository(db *sql.DB) domain.AuctionRepository {
	return &AuctionRepository{db: db}
}

func (r *AuctionRepository) GetByID(ctx context.Context, id int64) (*entities.Auction, error) {
	query := `
        SELECT id, lot_id, status, winner_id, final_price, auction_start, auction_end
        FROM auctions WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)
	auction := &entities.Auction{}
	err := row.Scan(&auction.ID, &auction.LotID, &auction.Status, &auction.WinnerID, &auction.FinalPrice, &auction.AuctionStart, &auction.AuctionEnd)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return auction, nil
}

func (r *AuctionRepository) Update(ctx context.Context, auction *entities.Auction) error {
	query := `
        UPDATE auctions SET status = $1, winner_id = $2, final_price = $3
        WHERE id = $4
    `
	_, err := r.db.ExecContext(ctx, query, auction.Status, auction.WinnerID, auction.FinalPrice, auction.ID)
	return err
}

func (r *AuctionRepository) GetAuctionsEndingBefore(ctx context.Context, endTime time.Time) ([]*entities.Auction, error) {
	query := `
        SELECT id, lot_id, status, winner_id, final_price, auction_start, auction_end
        FROM auctions WHERE auction_end <= $1 AND status = 'OPEN'
    `
	rows, err := r.db.QueryContext(ctx, query, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auctions []*entities.Auction
	for rows.Next() {
		auction := &entities.Auction{}
		err := rows.Scan(&auction.ID, &auction.LotID, &auction.Status, &auction.WinnerID, &auction.FinalPrice, &auction.AuctionStart, &auction.AuctionEnd)
		if err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}
	return auctions, nil
}
