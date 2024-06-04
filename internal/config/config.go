package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer struct {
		Address         string        `env:"HTTP_SERVER_ADDRESS" env-required:"true"`
		ShutdownTimeout time.Duration `env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" env-required:"true"`
	} `env-prefix:""`

	GRPCAuthServer struct {
		Address string `env:"GRPC_AUTH_SERVER_ADDRESS" env-required:"true"`
	} `env-prefix:""`

	GRPCTaskServer struct {
		Address string `env:"GRPC_TASK_SERVER_ADDRESS" env-required:"true"`
	}
}

func MustLoadConfig() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	return &cfg
}
