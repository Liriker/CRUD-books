package config

type Config struct {
	User     string
	Password string
}

func New() *Config {
	return &Config{
		User:     "root",
		Password: "21012001Ilya",
	}
	//	TODO normal config
}
