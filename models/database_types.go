package models

type Terminal struct {
	Id                    int
	Owner                 string
	Address               string
	Latitudes, Longitudes float64
	BunkerIn, BunkerOut   float64
	RateIn, RateOut       float64
}

type Team struct {
	Id      int
	Captain string
}

type Rasp struct {
	Id           int
	Team_id      int
	Terminal_id  int
	Terminal_num int
}
