package crosscutting

import (
	"github.com/kelseyhightower/envconfig"
)

type dbProvider string

const (
	MongoProvider    dbProvider = "mongo"
	PostgresProvider dbProvider = "postgres"
)

var AppSettings appSettings

type appSettings struct {
	DbSetting dbSetting
}

type dbSetting struct {
	Provider dbProvider `envconfig:"DB_PROVIDER" required:"true"`
	Host     string     `envconfig:"DB_HOST" required:"true"`
	Port     int        `envconfig:"DB_PORT" required:"true"`
	User     string     `envconfig:"DB_USER"`
	Password string     `envconfig:"DB_PASS"`
}

func init() {
	if err := envconfig.Process("", &AppSettings); err != nil {
		panic(err.Error())
	}
}
