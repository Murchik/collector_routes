package database

import (
	"context"
	"log"

	"github.com/Murchik/collector_routes/models"
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

func SelectQuery(conn *pgx.Conn) {

	// https://pkg.go.dev/github.com/jackc/pgx#hdr-Query_Interface
	// https://www.sohamkamani.com/golang/sql-database/
	qry := "SELECT * FROM terminals"

	rows, err := conn.Query(context.Background(), qry)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	terminals := []models.Terminal{}

	for rows.Next() {
		terminal := models.Terminal{}

		err := rows.Scan(&terminal.Id, &terminal.Owner, &terminal.Address, &terminal.Latitudes, &terminal.Longitudes)
		if err != nil {
			log.Fatal(err)
		}

		terminals = append(terminals, terminal)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}
}

func InsertQuery(conn *pgx.Conn) {

	commandTag, err := conn.Exec(context.Background(), "DELETE FROM widgets WEHRE id=$1", 42)
	if err != nil {
		log.Fatal(err)
	}

	if commandTag.RowsAffected() != 1 {
		log.Println("No row found to delete")
	}
}
