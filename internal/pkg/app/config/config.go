package config

import (
	"CRUD-books/internal/pkg/app/config/db"
	"CRUD-books/internal/pkg/app/config/server"
	"strconv"
)

type Config struct {
	server *server.Server `mapstructure:"server"`
	mySql  *db.DB
}

func New(path, filename string) (*Config, error) {
	serv, err := server.New(path, filename)
	if err != nil {
		return nil, err
	}

	mySql, err := db.New()
	if err != nil {
		return nil, err
	}

	return &Config{
		server: serv,
		mySql:  mySql,
	}, nil
}

func (c *Config) UserAndPassword() (string, string) {
	return c.mySql.User(), c.mySql.Password()
}

func (c *Config) Port() string {
	port := strconv.FormatInt(int64(c.server.Port), 10)
	return ":" + port
}

func (c *Config) Host() string {
	return c.server.Host
}
