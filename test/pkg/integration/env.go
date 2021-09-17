package integration

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Host string `env:"IDEMPOTENT_REQUESTS_HOST" envDefault:"http://localhost:8080"`
}

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
