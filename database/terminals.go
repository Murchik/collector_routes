package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Murchik/collector_routes/models"
	"github.com/jackc/pgx/v5"
)

const (
	terminalColumns = "id, owner, address, latitudes, longitudes, bunker_in, bunker_out, rate_in, rate_out"
	insertValues    = "$1, $2, $3, $4, $5, $6, $7, $8, $9"
)

func InsertTerminal(conn *pgx.Conn, t models.Terminal) {
	sqlStr := fmt.Sprintf("INSERT INTO terminals (%s) VALUES(%s)", terminalColumns, insertValues)

	commandTag, err := conn.Exec(context.Background(), sqlStr,
		t.Id,
		t.Owner,
		t.Address,
		t.Latitudes,
		t.Longitudes,
		t.BunkerIn,
		t.BunkerOut,
		t.RateIn,
		t.RateOut)

	if err != nil {
		log.Println(err.Error())
	}

	if commandTag.RowsAffected() != 1 {
		log.Println("No rows were inserted")
	}
}

func InsertTerminals(conn *pgx.Conn, terminals []models.Terminal) {
	batch := &pgx.Batch{}

	sqlStr := fmt.Sprintf("INSERT INTO terminals (%s) VALUES(%s)", terminalColumns, insertValues)

	for _, t := range terminals {
		batch.Queue(sqlStr,
			t.Id,
			t.Owner,
			t.Address,
			t.Latitudes,
			t.Longitudes,
			t.BunkerIn,
			t.BunkerOut,
			t.RateIn,
			t.RateOut)
	}

	batchResult := conn.SendBatch(context.Background(), batch)
	defer batchResult.Close()

	commandTag, err := batchResult.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}

	if commandTag.RowsAffected() != 1 {
		log.Println("No rows were inserted")
	}
}

func SelectTerminals(conn *pgx.Conn) []models.Terminal {
	var terminals []models.Terminal

	rows, err := conn.Query(context.Background(), "SELECT * FROM terminals")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Terminal

		err := rows.Scan(&t.Id, &t.Owner, &t.Address, &t.Latitudes, &t.Longitudes, &t.BunkerIn, &t.BunkerOut, &t.RateIn, &t.RateOut)
		if err != nil {
			log.Println(err.Error())
		}

		terminals = append(terminals, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return terminals
}
