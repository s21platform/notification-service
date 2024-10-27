package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type key string

const KeyMetrics = key("metrics")
const KeyUUID = key("uuid")

type Config struct {
	Service           Service
	Postgres          Postgres
	User              User
	EmailVerification EmailVerification
	Kafka             Kafka
	Platform          Platform
	Metrics           Metrics
}

type Service struct {
	Port string `env:"NOTIFICATION_SERVICE_PORT"`
}

type Postgres struct {
	User     string `env:"NOTIFICATION_SERVICE_POSTGRES_USER"`
	Password string `env:"NOTIFICATION_SERVICE_POSTGRES_PASSWORD"`
	Database string `env:"NOTIFICATION_SERVICE_POSTGRES_DB"`
	Host     string `env:"NOTIFICATION_SERVICE_POSTGRES_HOST"`
	Port     string `env:"NOTIFICATION_SERVICE_POSTGRES_PORT"`
}

type User struct {
	Host string `env:"USER_SERVICE_HOST"`
	Port string `env:"USER_SERVICE_PORT"`
}

type EmailVerification struct {
	Server   string `env:"EMAIL_SERVER"`
	Port     int    `env:"EMAIL_PORT"`
	User     string `env:"EMAIL_VERIFICATION_USER"`
	Password string `env:"EMAIL_VERIFICATION_PASSWORD"`
}

type Kafka struct {
	NotificationNewFriendTopic string `env:"FRIENDS_EMAIL_INVITE"`
	Server                     string `env:"KAFKA_SERVER"`
	GroupID                    string `env:"KAFKA_GROUP_ID"`
	AutoOffset                 string `env:"KAFKA_OFFSET"`
}

type Metrics struct {
	Host string `env:"GRAFANA_HOST"`
	Port int    `env:"GRAFANA_PORT"`
}

type Platform struct {
	Env string `env:"ENV"`
}

func MustLoad() *Config {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("failed to read env variables: %s", err)
	}
	return cfg
}
