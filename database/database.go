package database

import (
	"context"
	"fmt"
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

func ClearTable(conn *pgx.Conn, tableName string) {
	sqlStr := fmt.Sprintf("DELETE FROM %s;", tableName)

	commandTag, err := conn.Exec(context.Background(), sqlStr)
	if err != nil {
		log.Println(err.Error())
	}

	if commandTag.RowsAffected() == 0 {
		log.Println("No rows were deleted")
	}
}
