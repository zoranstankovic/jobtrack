-- name: Ping :one
-- Ping checks that the database answers a query. Used by /healthz.
SELECT 1::int AS ok;
