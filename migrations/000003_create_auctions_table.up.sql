CREATE TABLE IF NOT EXISTS auctions (
                                        id SERIAL PRIMARY KEY,
                                        lot_id INT NOT NULL REFERENCES lots(id),
    status VARCHAR(10) NOT NULL DEFAULT 'OPEN',
    winner_id INT REFERENCES users(id),
    final_price NUMERIC(10,2),
    auction_start TIMESTAMP NOT NULL,
    auction_end TIMESTAMP NOT NULL
    );
