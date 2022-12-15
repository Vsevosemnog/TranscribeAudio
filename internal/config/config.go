package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Conf struct {
	ENV      string `envconfig:"ENV" default:"dev"`
	DB       backoffice
	RabbitMQ rabbitmq
	TG       telegram
	Zoom     zoom
	AWS      aws
	YC       yc
}

type backoffice struct {
	URL string `envconfig:"PGSQL_BACKOFFICE_URL" default:"postgres://postgres:postgres@127.0.0.1/db_name"`
}

type telegram struct {
	BotApiKey string `envconfig:"TELEGRAM_BOT_API_KEY" required:"true"`
	ChId      int64  `envconfig:"TELEGRAM_CHAT_ID" required:"true"`
	LogDebug  bool   `envconfig:"TELEGRAM_BOT_DEBUG" default:"false"`
}

type zoom struct {
	ApiKey    string `envconfig:"ZOOM_API_KEY" required:"true"`
	ApiSecret string `envconfig:"ZOOM_API_SECRET" required:"true"`
}

type aws struct {
	KeyID     string `envconfig:"AWS_ACCESS_KEY_ID"`
	KeySecret string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	Region    string `envconfig:"AWS_DEFAULT_REGION"`
}

type yc struct {
	StaticKeyID     string `envconfig:"YC_STATIC_ACCESS_KEY_ID"`
	StaticKeySecret string `envconfig:"YC_STATIC_SECRET_ACCESS_KEY"`
	Region          string `envconfig:"YC_DEFAULT_REGION" default:"ru-central1"`
	IamSaKey        []byte `envconfig:"YC_IAM_SA_KEY"`
}

type rabbitmq struct {
	SSL            bool     `envconfig:"RABBITMQ_SSL" default:"false"`
	Host           string   `envconfig:"RABBITMQ_HOST" default:"127.0.0.1"`
	Port           string   `envconfig:"RABBITMQ_PORT" default:"5672"`
	Username       string   `envconfig:"RABBITMQ_USER" default:"guest"`
	Password       string   `envconfig:"RABBITMQ_PASSWORD" default:"guest"`
	StructPrefixes []string `envconfig:"RABBITMQ_STRUCT_PREFIXES" default:"UserPayments"`
}

func Load() *Conf {
	var c Conf

	err := envconfig.Process("envconfig", &c)
	if err != nil {
		fmt.Printf("can't read env: %s", err)
	}
	return &c
}
