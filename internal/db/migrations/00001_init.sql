-- Initial migration. It is intentionally empty: it only proves that goose
-- runs at startup and records version 1 in the goose_db_version table.
-- Real tables arrive with the next tickets, each in a new numbered file
-- (00002_companies.sql, ...). Never edit a migration that has already run;
-- add a new one instead.

-- +goose Up

-- +goose Down
