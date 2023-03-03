package config

import (
	"CRUD-books/internal/pkg/app/config/db"
	"CRUD-books/internal/pkg/app/config/server"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"strconv"
)

type Config struct {
	Server *server.Server `mapstructure:"Server"`
	MySql  *db.DB         `mapstructure:"Db"`
	logs   zerolog.Logger
}

func New(filename, confPath string) (*Config, error) {

	conf := new(Config)
	viper.AddConfigPath(confPath)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	//if conf.MySql.User == "" {
	//	err := errors.New("have no user in DB config")
	//	return nil, err
	//}

	return conf, nil
}

func (c *Config) User() string {
	return c.MySql.User
}

func (c *Config) Password() string {
	return c.MySql.Password
}

func (c *Config) Port() string {
	port := strconv.FormatInt(int64(c.Server.Port), 10)
	return ":" + port
}

func (c *Config) Host() string {
	return c.Server.Host
}
