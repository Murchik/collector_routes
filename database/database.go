package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context, url string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
