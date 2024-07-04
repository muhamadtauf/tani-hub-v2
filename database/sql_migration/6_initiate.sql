-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    sub_title TEXT NOT NULL,
    content TEXT NOT NULL,
    is_at_home BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd