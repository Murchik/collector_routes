package main

import (
	"encoding/xml"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/Murchik/collector_routes/packages/atm"
	"github.com/Murchik/collector_routes/packages/pathfinding"
)

func main() {
	// Получить ATMs в структурку
	log.Println("Making request...")
	atms, err := atm.GetATMs()
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

	// Создать рандомные пути
	log.Println("Making matrix...")
	const qnt int = 1000
	arr := make([][]float64, qnt)
	for i := range arr {
		arr[i] = make([]float64, qnt)
	}
	for i := 0; i < qnt; i++ {
		for j := 0; j < qnt; j++ {
			if i == j {
				continue
			}
			arr[i][j] = rand.Float64()*30 + 30
		}
	}

	// банкоматы на завтра
	atms_1 := atm.GetAtmsOnDay(atms, 1)

	f, err := os.Create("./routes.txt")

	for group := 0; group < 5; group++ {

		// Найти путь
		log.Println("Searching for path...")
		res := pathfinding.Pathfinding(atms_1[0:qnt], arr, atms_1[0]) // сейчас кдаляется начальный банкомат, надо это исправить и сделать начальный банкомат всегда одинаковый либо заменить его на какую то другую точку
		// Удалить найденные банкоматы из общего массива
		for i := 0; i < len(atms_1); i++ {
			for j := 0; j < len(res); j++ {
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
	}

}
