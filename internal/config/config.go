package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// type Config struct {
// 	HTTPServer      HTTPServer      `yaml:"http_server"`
// 	GrpcAuthServer  GrpcAuthServer  `yaml:"grpc_auth_server"`
// 	GrpcTasksServer GrpcTasksServer `yaml:"grpc_tasks_server"`
// }

// type HTTPServer struct {
// 	Address         string        `yaml:"address"`
// 	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
// }

// type GrpcAuthServer struct {
// 	Address string `yaml:"address"`
// }

// type GrpcTasksServer struct {
// 	Address string `yaml:"address"`
// }

// func MustLoadConfig(configPath string) *Config {
// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		panic("config file does not exists: " + configPath)
// 	}

// 	var cfg Config

// 	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
// 		panic("cannot read config: " + err.Error())
// 	}

// 	return &cfg
// }

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
