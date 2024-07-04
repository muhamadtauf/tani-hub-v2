-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    price DECIMAL NOT NULL,
    quantity INTEGER NOT NULL,
    product_id BIGINT NOT NULL,
    order_id BIGINT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (order_id) REFERENCES orders (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd