package models

import (
	"log"
	"math/rand"

	"github.com/Murchik/collector_routes/overpass"
)

// Возвращает случайное число от 0 до 1
func getRandomNumber() float64 {
	return rand.Float64()
}

// Возвращает массив с банкоматами
func GetATMs() ([]Terminal, error) {
	moscow := overpass.City{Name: "Moscow", Radius: 18000.0, Lat: 55.752221, Lon: 37.623978}
	data, err := overpass.MakeQuery(moscow, "atm")
	if err != nil {
		return nil, err
	}

	var atms []Terminal
	for i, v := range data.Nodes {
		atms = append(atms, Terminal{
			Id:         i,
			Owner:      "",
			Address:    "",
			Latitudes:  v.Latitude,
			Longitudes: v.Longitude,
			BunkerIn:   getRandomNumber(),
			BunkerOut:  getRandomNumber(),
			RateIn:     getRandomNumber(),
			RateOut:    getRandomNumber(),
		})
	}
	return atms, nil
}

func GetAtmsOnDay(atms []Terminal, day int) []Terminal {
	var res []Terminal

	log.Println("Searching for ATMs which need to visit on day ", day, "...")
	for _, atm := range atms {
		if (atm.BunkerIn-float64(day)*atm.RateIn <= 0) || (atm.BunkerOut-float64(day)*atm.RateOut <= 0) {
			res = append(res, atm)
		}
	}

	return res
}
