package models

type Terminal struct {
	Id                                            int
	Owner                                         string
	Address                                       string
	Latitudes, Longitudes                         float64
	Bunker_priema_procent, Bunker_vidachi_procent int
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
