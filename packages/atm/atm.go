package atm

import (
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
		atms = append(atms, ATM{i, v.Latitude, v.Longitude, 0.3, 0.3, 0.25, 0.25})
	}
	return atms, nil
}
