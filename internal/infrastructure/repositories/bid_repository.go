// internal/infrastructure/repositories/bid_repository.go

package repositories

import (
	"context"
	"database/sql"

	"indiv/internal/domain/entities"
	domain "indiv/internal/domain/repositories"
)

type BidRepository struct {
	db *sql.DB
}

func NewBidRepository(db *sql.DB) domain.BidRepository {
	return &BidRepository{db: db}
}

func (r *BidRepository) Create(ctx context.Context, bid *entities.Bid) error {
	query := `
        INSERT INTO bids (auction_id, bidder_id, amount, timestamp)
        VALUES ($1, $2, $3, $4) RETURNING id
    `
	return r.db.QueryRowContext(ctx, query, bid.AuctionID, bid.BidderID, bid.Amount, bid.Timestamp).Scan(&bid.ID)
}

func (r *BidRepository) GetByAuctionID(ctx context.Context, auctionID int64) ([]*entities.Bid, error) {
	query := `
        SELECT id, auction_id, bidder_id, amount, timestamp
        FROM bids WHERE auction_id = $1 ORDER BY amount DESC
    `
	rows, err := r.db.QueryContext(ctx, query, auctionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bids []*entities.Bid
	for rows.Next() {
		bid := &entities.Bid{}
		err := rows.Scan(&bid.ID, &bid.AuctionID, &bid.BidderID, &bid.Amount, &bid.Timestamp)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}
	return bids, nil
}

func (r *BidRepository) GetHighestBid(ctx context.Context, auctionID int64) (*entities.Bid, error) {
	query := `
        SELECT id, auction_id, bidder_id, amount, timestamp
        FROM bids WHERE auction_id = $1 ORDER BY amount DESC LIMIT 1
    `
	row := r.db.QueryRowContext(ctx, query, auctionID)
	bid := &entities.Bid{}
	err := row.Scan(&bid.ID, &bid.AuctionID, &bid.BidderID, &bid.Amount, &bid.Timestamp)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return bid, nil
}

func (r *BidRepository) RefundLosingBids(ctx context.Context, auctionID, winnerID int64) error {
	query := `
        SELECT bidder_id, amount FROM bids
        WHERE auction_id = $1 AND bidder_id != $2
    `
	rows, err := r.db.QueryContext(ctx, query, auctionID, winnerID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var bidderID int64
		var amount float64
		if err := rows.Scan(&bidderID, &amount); err != nil {
			return err
		}

		// Возврат средств проигравшему участнику
		_, err := r.db.ExecContext(ctx, `
            UPDATE users SET balance = balance + $1 WHERE id = $2
        `, amount, bidderID)
		if err != nil {
			return err
		}
	}
	return nil
}
