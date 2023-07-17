CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
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
, activity_pub_id TEXT);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE statuses
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER   NOT NULL,
    text       TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20230716124959'),
  ('20230716125041'),
  ('20230717150225');
