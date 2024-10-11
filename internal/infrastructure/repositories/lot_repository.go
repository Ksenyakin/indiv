// internal/infrastructure/repositories/lot_repository.go

package repositories

import (
	"context"
	"database/sql"

	"indiv/internal/domain/entities"
	domain "indiv/internal/domain/repositories"
)

type LotRepository struct {
	db *sql.DB
}

func NewLotRepository(db *sql.DB) domain.LotRepository {
	return &LotRepository{db: db}
}

func (r *LotRepository) Create(ctx context.Context, lot *entities.Lot) error {
	query := `
        INSERT INTO lots (seller_id, title, description, starting_price, min_bid_increment, auction_start, auction_end)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
    `
	return r.db.QueryRowContext(ctx, query, lot.SellerID, lot.Title, lot.Description, lot.StartingPrice, lot.MinBidIncrement, lot.AuctionStart, lot.AuctionEnd).Scan(&lot.ID)
}

func (r *LotRepository) GetByID(ctx context.Context, id int64) (*entities.Lot, error) {
	query := `
        SELECT id, seller_id, title, description, starting_price, min_bid_increment, auction_start, auction_end
        FROM lots WHERE id = $1
    `
	row := r.db.QueryRowContext(ctx, query, id)
	lot := &entities.Lot{}
	err := row.Scan(&lot.ID, &lot.SellerID, &lot.Title, &lot.Description, &lot.StartingPrice, &lot.MinBidIncrement, &lot.AuctionStart, &lot.AuctionEnd)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return lot, nil
}

func (r *LotRepository) List(ctx context.Context, page, pageSize int32) ([]*entities.Lot, error) {
	offset := (page - 1) * pageSize
	query := `
        SELECT id, seller_id, title, description, starting_price, min_bid_increment, auction_start, auction_end
        FROM lots ORDER BY id LIMIT $1 OFFSET $2
    `
	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lots []*entities.Lot
	for rows.Next() {
		lot := &entities.Lot{}
		err := rows.Scan(&lot.ID, &lot.SellerID, &lot.Title, &lot.Description, &lot.StartingPrice, &lot.MinBidIncrement, &lot.AuctionStart, &lot.AuctionEnd)
		if err != nil {
			return nil, err
		}
		lots = append(lots, lot)
	}
	return lots, nil
}
