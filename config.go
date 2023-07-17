package shirase

import (
	"fmt"
	"net/url"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: DIしてぇ〜〜〜〜〜!!!!!!!!
var GlobalConfig Config

type Config struct {
	Bind        string
	Port        int    `default:"3000"`
	LocalDomain string `required:"true" split_words:"true"`
	Insecure    bool   `default:"false"`
}

func (c *Config) HttpListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}

func (c *Config) URLBase() *url.URL {
	scheme := "https://"
	if c.Insecure {
		scheme = "http://"
	}
	u, err := url.Parse(scheme + c.LocalDomain)
	if err != nil {
		panic(err)
	}
	return u
}
