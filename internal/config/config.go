package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
	"time"
)

type httpConfig struct {
	Ip   string
	Port string
}

type Database struct {
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type Config struct {
	Server   httpConfig
	Database Database
	ApiKey   string
}

func New() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server: httpConfig{
			Ip:   os.Getenv("APP_IP"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: Database{
			Dsn:             os.Getenv("DB_DSN"),
			MaxIdleConns:    cast.ToInt(os.Getenv("DB_MAX_IDLE_CONNS")),
			MaxOpenConns:    cast.ToInt(os.Getenv("DB_MAX_OPEN_CONNS")),
			ConnMaxLifetime: time.Duration(cast.ToInt(os.Getenv("DB_CONN_MAX_LIFETIME_MINUTES"))) * time.Minute,
		},
		ApiKey: os.Getenv("APP_KEY"),
	}, nil
}
