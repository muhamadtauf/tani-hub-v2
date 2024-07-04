-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    price DECIMAL NOT NULL,
    quantity INTEGER NOT NULL,
    total DECIMAL NOT NULL,
    product_id BIGINT NOT NULL,
    order_uuid VARCHAR(256) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (order_uuid) REFERENCES orders (uuid),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd