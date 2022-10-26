package atm

type coordinates struct {
	lat float64
	lon float64
}

type ATM struct {
	id                    int
	coordinates           coordinates
	bunker_in, bunker_out float64
	rate_in, rate_out     float64
}
