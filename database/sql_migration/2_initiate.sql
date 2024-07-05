-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    code VARCHAR(256) UNIQUE NOT NULL,
    status VARCHAR(256) NOT NULL,
    total DECIMAL NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +migrate StatementEnd