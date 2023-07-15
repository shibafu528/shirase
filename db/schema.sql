CREATE TABLE accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    domain TEXT,
    display_name TEXT,
    private_key TEXT,
    public_key TEXT
);
