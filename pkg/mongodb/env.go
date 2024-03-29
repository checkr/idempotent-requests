package mongodb

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	URI        string `env:"MONGODB_URI" envDefault:"mongodb://root:password123@localhost:27017"`
	Database   string `env:"MONGODB_DATABASE" envDefault:"idempotent-requests"`
	Collection string `env:"MONGODB_COLLECTION" envDefault:"captures"`
}

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
