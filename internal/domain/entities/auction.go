// internal/domain/entities/auction.go

package entities

import "time"

type Auction struct {
	ID           int64     `json:"id"`
	LotID        int64     `json:"lot_id"`
	Status       string    `json:"status"` // "OPEN" или "CLOSED"
	WinnerID     *int64    `json:"winner_id,omitempty"`
	FinalPrice   *float64  `json:"final_price,omitempty"`
	AuctionStart time.Time `json:"auction_start"`
	AuctionEnd   time.Time `json:"auction_end"`
}
