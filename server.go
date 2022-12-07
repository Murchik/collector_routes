package main

import (
	"context"
	"encoding/xml"
	"log"
	"os"

	"github.com/Murchik/collector_routes/config"
	"github.com/Murchik/collector_routes/database"
	"github.com/Murchik/collector_routes/models"
)

func main() {
	// Подключение к базе данных
	ctx := context.Background()
	db := database.CreateConnection(ctx, config.GetURL())
	defer db.Close(ctx)

	// Получение всех терминалов из базы данных
	//terminals := database.SelectTerminals(db)
	//log.Println(terminals)

	// Добавление в базу данных одного терминала
	//database.InsertTerminal(db, models.Terminal{Id: 104})

	// Добавление в базу данных массива терминалов
	// terminals = []models.Terminal{
	// 	{Id: 104},
	// 	{Id: 105},
	// 	{Id: 106},
	// }
	// database.InsertTerminals(db, terminals)

	// log.Fatal("AfterDatabaseConnect")

	// Получить ATMs в структурку
	log.Println("Making request...")
	atms, err := models.GetATMs()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Записать структурку в файл
	log.Println("Writing into osmOutput.xml...")
	xml, err := xml.MarshalIndent(atms, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = os.WriteFile("osmOutput.xml", xml, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
