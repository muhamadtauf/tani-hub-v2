-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd