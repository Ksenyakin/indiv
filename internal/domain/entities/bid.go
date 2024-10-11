// internal/domain/entities/bid.go

package entities

import "time"

type Bid struct {
	ID        int64     `json:"id"`
	AuctionID int64     `json:"auction_id"`
	BidderID  int64     `json:"bidder_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
