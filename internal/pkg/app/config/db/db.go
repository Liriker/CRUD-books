package db

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
