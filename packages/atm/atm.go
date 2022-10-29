package atm

type Coordinates struct {
	Lat float64
	Lon float64
}

type ATM struct {
	Id                    int
	Coordinates           Coordinates
	Bunker_in, Bunker_out float64
	Rate_in, Rate_out     float64
}
