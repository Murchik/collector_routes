package atm

import (
	"log"
	"strconv"

	"github.com/Murchik/collector_routes/packages/overpass"
)

type ATM struct {
	Id                  int
	Latitude, Longitude float64
	BunkerIn, BunkerOut float64
	RateIn, RateOut     float64
}

func GetATMs() ([]ATM, error) {
	moscow := overpass.City{Name: "Moscow", Radius: 18000.0, Lat: 55.752221, Lon: 37.623978}
	data, err := overpass.MakeQuery(moscow, "atm")
	if err != nil {
		return nil, err
	}

	var atms []ATM
	for i, v := range data.Nodes {
		atms = append(atms, ATM{i, v.Latitude, v.Longitude, 0.3, 0.3, 0.7, 0.25}) // Напихать рандомные данные
	}
	return atms, nil
}

func GetAtmsOnDay(atms []ATM, day int64) []ATM {

	var res []ATM

	log.Println("Searching for ATMs which need to visit on day " + strconv.FormatInt(day, 10) + "...")

	for _, atm := range atms {
		if (atm.BunkerIn-float64(day)*atm.RateIn <= 0) || (atm.BunkerOut-float64(day)*atm.RateOut <= 0) {
			res = append(res, atm)
		}
	}

	return res
}
