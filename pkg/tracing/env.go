package tracing

import (
	"github.com/caarlos0/env/v6"
)

const Datadog = "datadog"

type Config struct {
	TracingEnabled       bool   `env:"OPENTRACING_ENABLED" envDefault:"true"`
	TracerImplementation string `env:"OPENTRACING_TRACER_IMPLEMENTATION" envDefault:"datadog"`
}

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
