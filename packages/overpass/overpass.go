package overpass

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Node struct {
	Id  int     `xml:"id,attr"`
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

type ATMS struct {
	XMLName xml.Name
	Atms    []Node
}

func GetATMs() {
	//55.66076 / 37.480117 (lat/lon)
	//55.809556 / 37.7053288 (lat/lon)

	OPserv := "https://overpass-api.de/api/interpreter?data="
	OPdata := "node[amenity=atm](55.66076,37.480117,55.809556,37.7053288);out;"

	querry := fmt.Sprintf("%s%s", OPserv, OPdata)
	fmt.Println(querry)

	resp, err := http.Get(querry)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	// open input file
	fo, err := os.Create("input.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	var atms ATMS

	xml.Unmarshal(body, &atms)

	if _, err := fo.Write(body); err != nil {
		panic(err)
	}

	fmt.Println(atms)
	//fmt.Println(atms.XMLName)
	//fmt.Println(atms.Atms[0])

}

func MakeQuerry() {

	resp, err := http.Get("https://overpass-api.de/api/interpreter?data=node(1);out;")

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}
