CREATE TABLE IF NOT EXISTS market_listing(
    id serial PRIMARY KEY,
    code VARCHAR (50) NOT NULL,
    order_type VARCHAR (15) NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    amount BIGINT NOT NULL,
    listing_date TIMESTAMP NOT NULL,
    expiry_date TIMESTAMP NOT NULL,
    CONSTRAINT fk_code
        FOREIGN KEY (code)
            REFERENCES entity(code)
);