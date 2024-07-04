-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    code VARCHAR(256) UNIQUE NOT NULL,
    price DECIMAL NOT NULL,
    stock INTEGER NOT NULL,
    is_at_home BOOLEAN NOT NULL,
    category_id BIGINT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd