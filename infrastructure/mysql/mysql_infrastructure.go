package mysql

import (
	"context"
	"database/sql"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"time"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Database.Dsn)
	if err != nil {
		return nil, err
	}

	if config.Database.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.Database.MaxIdleConns)
	}
	if config.Database.MaxOpenConns > 0 {
		db.SetMaxOpenConns(config.Database.MaxOpenConns)
	}
	if config.Database.ConnMaxLifetime.Nanoseconds() > 0 {
		db.SetConnMaxLifetime(config.Database.ConnMaxLifetime)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
