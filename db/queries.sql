-- name: GetAccount :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts WHERE username = ? LIMIT 1;

-- name: CreateAccount :execresult
INSERT INTO accounts (
    username
) VALUES (
    ?
);

-- name: CreateStatus :execresult
INSERT INTO statuses (
    account_id, text
) VALUES(
    ?, ?
)
