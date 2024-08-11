package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service  Service
	Postgres ReadEnvDB
}
type Service struct {
	Port string `env:"NOTIFICATION_SERVICE_PORT"`
}

type ReadEnvDB struct {
	User     string `env:"NOTIFICATION_SERVICE_POSTGRES_USER"`
	Password string `env:"NOTIFICATION_SERVICE_POSTGRES_PASSWORD"`
	Database string `env:"NOTIFICATION_SERVICE_POSTGRES_DB"`
	Host     string `env:"NOTIFICATION_SERVICE_POSTGRES_HOST"`
	Port     string `env:"NOTIFICATION_SERVICE_POSTGRES_PORT"`
}

func MustLoad() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Can not read env variables: %s", err)
	}
	return cfg
}
