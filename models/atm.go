package models

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/Murchik/collector_routes/overpass"
)

// Возвращает случайное число от 0 до 1
func getRandomNumber() float64 {
	return rand.Float64()
}

type ATM struct {
	Id                  int     // id банкомата
	Latitude, Longitude float64 // Координаты
	BunkerIn, BunkerOut float64 // Состояние бункеров
	RateIn, RateOut     float64 // Скорость заполнения бункеров
}

// Возвращает массив с банкоматами
func GetATMs() ([]ATM, error) {
	moscow := overpass.City{Name: "Moscow", Radius: 18000.0, Lat: 55.752221, Lon: 37.623978} // Диапазон поиска банкоматов
	data, err := overpass.MakeQuery(moscow, "atm")                                           // Запрос к overpass
	if err != nil {
		return nil, err
	}

	var atms []ATM
	for i, v := range data.Nodes {
		atms = append(atms, ATM{i, v.Latitude, v.Longitude, getRandomNumber(), getRandomNumber(), getRandomNumber(), getRandomNumber()})
	}
	return atms, nil
}

// Входные данные - массив банкоматов, день на который надо посетить банкомат
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
