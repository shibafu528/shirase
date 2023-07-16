-- migrate:up
CREATE TABLE statuses
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER   NOT NULL,
    text       TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE statuses;
