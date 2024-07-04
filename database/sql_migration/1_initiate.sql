-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(256) UNIQUE NOT NULL,
    name VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    role VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd