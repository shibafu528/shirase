-- name: GetAccount :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts WHERE username = ? LIMIT 1;

-- name: GetAccountByActivityPubID :one
SELECT * FROM accounts WHERE activity_pub_id = ? LIMIT 1;

-- name: GetAccountIDByActivityPubID :one
SELECT id FROM accounts WHERE activity_pub_id = ? LIMIT 1;

-- name: GetStatusesByAccountID :many
SELECT
    s.id,
    s.account_id,
    s.text,
    s.created_at,
    s.updated_at,
    coalesce(a.activity_pub_id, a.username) activity_pub_id
FROM statuses s INNER JOIN accounts a ON s.account_id = a.id WHERE s.account_id = ? ORDER BY s.id DESC;

-- name: CreateAccount :execresult
INSERT INTO accounts (
    username, activity_pub_id, private_key, public_key
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateStatus :execresult
INSERT INTO statuses (
    account_id, text
) VALUES(
    ?, ?
)
