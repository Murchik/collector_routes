package main

import (
	"context"
	"encoding/xml"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/Murchik/collector_routes/config"
	"github.com/Murchik/collector_routes/database"
	"github.com/Murchik/collector_routes/models"
	"github.com/Murchik/collector_routes/pathfinding"
)

func main() {
	const qnt int = 5

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

	atms = atms[:qnt] // Делаем это потому что матрица на 1000

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

	// Создать рандомные пути
	log.Println("Making matrix...")

	arr := make([][]float64, qnt)
	for i := range arr {
		arr[i] = make([]float64, qnt)
	}
	//var dinst float64 = 10
	for i := 0; i < qnt; i++ {
		for j := 0; j < qnt; j++ {
			if i == j {
				arr[i][j] = 0
			} else {
				if j < i {
					arr[i][j] = arr[j][i]
				} else {
					arr[i][j] = rand.Float64()*30 + 30
				}
			}
			// Сделать так чтобы arr[i][j] == arr[j][i]

			//arr[i][j] = dinst
			//dinst += 1
		}
	}

	//res := pathfinding.Pathfinding(atms[0:qnt], arr, atms[0])
	for i := 0; i < len(arr); i++ {
		log.Println(arr[i])
	}
	//log.Fatal("AfterSomething")

	// банкоматы на завтра
	atms_1 := models.GetAtmsOnDay(atms, 1)

	f, err := os.Create("./routes.txt")

	for group := 0; group < 5; group++ {

		// Найти путь
		log.Println("Searching for path...")
		res := pathfinding.Pathfinding(atms_1[0:qnt], arr, atms_1[0])
		// Удалить найденные банкоматы из общего массива
		for i := 0; i < len(atms_1); i++ {
			for j := 1; j < len(res)-1; j++ {
				if atms_1[i].Id == res[j] {
					atms_1 = append(atms_1[:i], atms_1[i+1:]...)
					i -= 1
					break
				}
			}
		}

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
		break
	}

}
