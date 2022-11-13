package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"math/rand"
	"os"

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
	var arr [qnt][qnt]float64
	for i := 0; i < qnt; i++ {
		for j := 0; j < qnt; j++ {
			if i == j {
				continue
			}
			arr[i][j] = rand.Float64()*30 + 30
		}
	}

	// Найти путь
	log.Println("Searching for path...")
	res := pathfinding.Pathfinding(atms[0:qnt], arr, atms[0])

	// Записать путь в файл
	log.Println("Write results...")
	for index, id := range res {
		fmt.Println(index, id)
	}

}
