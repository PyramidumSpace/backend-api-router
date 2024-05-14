package env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	ConfigPath string
}

func MustLoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		panic("cannot load environment variables: " + err.Error())
	}

	configPath := mustGetEnv("CONFIG_PATH")

	return &Env{
		ConfigPath: configPath,
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is not set")
	}
	return value
}
