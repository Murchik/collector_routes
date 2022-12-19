package main

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"strconv"

	"github.com/Murchik/collector_routes/config"
	"github.com/Murchik/collector_routes/database"
	"github.com/Murchik/collector_routes/models"
	pf "github.com/Murchik/collector_routes/pathfinding"
)

func main() {
	// Подключение к базе данных
	ctx := context.Background()
	db := database.CreateConnection(ctx, config.GetURL())
	defer db.Close(ctx)

	// Очиститиь таблицу терминалов
	log.Println("Clearing terminals table...")
	database.ClearTable(db, "terminals")

	// Получить ATMs в структурку
	log.Println("Making request...")
	atms, err := models.GetATMs()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Сохраняем ATMs в базу данных
	log.Println("Inserting ATMs into db...")
	database.InsertTerminals(db, atms)

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

	const qnt int = 20
	atms = atms[:qnt]

	// Создать рандомные пути
	log.Println("Making distance matrix...")
	dist_matrix := pf.CreateDistanceMatrix(qnt)

	// банкоматы на завтра
	// atms_1 := models.GetAtmsOnDay(atms, 1)
	atms_1 := atms

	f, err := os.Create("./routes.txt")

	for group := 0; group < 5; group++ {

		// Найти путь
		log.Println("Searching for path for group " + strconv.FormatInt(int64(group), 10))
		res := pf.Pathfinding(atms_1, dist_matrix, atms_1[0])
		// Удалить найденные банкоматы из общего массива
		atms_1 = pf.DeleteAtmsFromArray(atms_1, res)

		// Записать путь в файл
		log.Println("Write results...")
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(res); i++ {
			f.WriteString(strconv.FormatInt(int64(res[i]), 10))
			f.WriteString(" ")
		}
		f.WriteString("\n")

		if len(atms_1) == 0 {
			break
		}
	}

}
