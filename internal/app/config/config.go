package config

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"0.0.0.0:80"`
	Flag          string `env:"FLAG"`
}
