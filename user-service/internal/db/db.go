package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}

	log.Println("Connected to Postgres")
	return pool
}
