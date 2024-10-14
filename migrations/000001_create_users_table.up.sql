CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
    balance NUMERIC(10,2) NOT NULL DEFAULT 0
    );
