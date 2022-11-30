package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnection() *pgx.Conn {

	db_url := "postgresql://maximov:123456@185.225.34.182:5432/test_db"

	config, err := pgx.ParseConfig(os.Getenv(db_url))
	if err != nil {
		log.Fatal("Cannot parse database url")
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	return conn

}
