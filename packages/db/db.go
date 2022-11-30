package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateConnection() *pgx.Conn {

	db_url := "postgresql://maximov:123456@185.225.34.182:5432/test_db"

	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		log.Fatal(err)
	}

	return conn

}
