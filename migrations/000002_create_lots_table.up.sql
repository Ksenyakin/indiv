CREATE TABLE IF NOT EXISTS lots (
                                    id SERIAL PRIMARY KEY,
                                    seller_id INT NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    starting_price NUMERIC(10,2) NOT NULL,
    min_bid_increment NUMERIC(10,2) NOT NULL,
    auction_start TIMESTAMP NOT NULL,
    auction_end TIMESTAMP NOT NULL
    );
