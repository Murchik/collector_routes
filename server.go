package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/Murchik/collector_routes/packages/atm"
	"github.com/Murchik/collector_routes/packages/overpass"
	"github.com/Murchik/collector_routes/packages/pathfinding"
)

func main() {
	const qnt int = 1000
	var atms []atm.ATM
	var arr [qnt][qnt]float64

	overpass_atms := overpass.GetATMs()

	log.Println("Making array of ATMs...")
	for index, atm_ := range overpass_atms {
		atms = append(atms, atm.ATM{Id: index, Latitude: atm_.Latitude, Longitude: atm_.Longitude, Bunker_in: 0.3, Bunker_out: 0.3, Rate_in: 0.25, Rate_out: 0.25})
	}

	log.Println("Making matrix...")
	for i := 0; i < qnt; i++ {
		for j := 0; j < qnt; j++ {
			if i == j {
				continue
			}
			arr[i][j] = rand.Float64()*30 + 30
		}
	}

	log.Println("Searching for path...")
	res := pathfinding.Pathfinding(atms, arr, atms[0])

	log.Println("Write results...")
	for index, id := range res {
		fmt.Println(index, id)
	}

}
