package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db      *sql.DB
	queries *Queries
)

// FIXME: こんな所に置くなバカ
func Open() (*sql.DB, *Queries) {
	if db == nil {
		d, err := sql.Open("sqlite3", "storage/database.db3") // TODO: 雑
		if err != nil {
			panic(err)
		}
		db = d
	}

	if queries == nil {
		queries = New(db)
	}

	return db, queries
}
