package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
	"time"
)

type http struct {
	ip   string
	port string
}

type Database struct {
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type Config struct {
	Server   http
	Database Database
}

func New() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server: http{
			ip:   os.Getenv("APP_IP"),
			port: os.Getenv("APP_PORT"),
		},
		Database: Database{
			Dsn:             os.Getenv("DB_DSN"),
			MaxIdleConns:    cast.ToInt(os.Getenv("DB_MAX_IDLE_CONNS")),
			MaxOpenConns:    cast.ToInt(os.Getenv("DB_MAX_OPEN_CONNS")),
			ConnMaxLifetime: time.Duration(cast.ToInt(os.Getenv("DB_CONN_MAX_LIFETIME_MINUTES"))) * time.Minute,
		},
	}, nil
}
