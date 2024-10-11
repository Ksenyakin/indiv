package entities

import "time"

type Lot struct {
	ID              int64     `json:"id"`
	SellerID        int64     `json:"seller_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	StartingPrice   float64   `json:"starting_price"`
	MinBidIncrement float64   `json:"min_bid_increment"`
	AuctionStart    time.Time `json:"auction_start"`
	AuctionEnd      time.Time `json:"auction_end"`
}
