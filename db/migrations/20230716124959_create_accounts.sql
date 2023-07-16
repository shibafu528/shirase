-- migrate:up
CREATE TABLE accounts
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    username     TEXT      NOT NULL,
    domain       TEXT,
    display_name TEXT,
    private_key  TEXT      NOT NULL,
    public_key   TEXT      NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE accounts;
