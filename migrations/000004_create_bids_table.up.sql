CREATE TABLE IF NOT EXISTS bids (
                                    id SERIAL PRIMARY KEY,
                                    auction_id INT NOT NULL REFERENCES auctions(id),
    bidder_id INT NOT NULL REFERENCES users(id),
    amount NUMERIC(10,2) NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW()
    );
