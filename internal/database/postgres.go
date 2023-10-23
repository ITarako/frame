package database

import (
	"context"
	"database/sql"
	"fmt"
	"frame/internal/config"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func NewPostgres(cfg *config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SslMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("sql open: %s", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		log.Fatalf("parse MaxIdleTime: %s", err)
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("ping contex: %s", err)
	}

	return db
}
