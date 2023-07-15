package shirase

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shibafu528/shirase/db"
)

// TODO: DIしてぇ〜〜〜〜〜!!!!!!!!
var GlobalConfig Config

type Config struct {
	Bind        string
	Port        int    `default:"3000"`
	LocalDomain string `required:"true" split_words:"true"`

	db      *sql.DB
	queries *db.Queries
}

func (c *Config) HttpListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}

func (c *Config) URLBase() *url.URL {
	u, err := url.Parse("http://" + c.LocalDomain)
	if err != nil {
		panic(err)
	}
	return u
}

func (c *Config) DB() (*sql.DB, *db.Queries) {
	if c.db == nil {
		d, err := sql.Open("sqlite3", "storage/database.db3") // TODO: 雑
		if err != nil {
			panic(err)
		}
		c.db = d
	}

	if c.queries == nil {
		c.queries = db.New(c.db)
	}

	return c.db, c.queries
}
