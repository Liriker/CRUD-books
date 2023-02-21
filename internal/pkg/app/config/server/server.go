package server

import "github.com/spf13/viper"

type Server struct {
	Host string `mapstructure:"Host"`
	Port int    `mapstructure:"Port"`
}

func New(confPath, filename string) (*Server, error) {
	serv := new(Server)
	viper.AddConfigPath(confPath)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&serv); err != nil {
		return nil, err
	}

	return serv, nil
}
