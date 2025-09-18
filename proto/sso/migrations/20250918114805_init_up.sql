-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS users
(
    id           INTEGER PRIMARY KEY,
    email        TEXT    NOT NULL UNIQUE,
    pass_hash    BLOB    NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id     INTEGER PRIMARY KEY,
    name   TEXT NOT NULL UNIQUE,
    secret TEXT NOT NULL UNIQUE
);

SELECT 'down SQL query';
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS apps;
-- +goose StatementBegin