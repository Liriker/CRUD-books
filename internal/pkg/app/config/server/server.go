package server

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
