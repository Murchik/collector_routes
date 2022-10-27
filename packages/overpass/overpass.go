package overpass

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
)

type Osm struct {
	XMLName   xml.Name `xml:"osm"`
	Version   string   `xml:"version,attr"`
	Generator string   `xml:"generator,attr"`
	Note      string   `xml:"note"`
	Meta      Meta     `xml:"meta"`
	Nodes     []Node   `xml:"node"`
}

type Meta struct {
	XMLName xml.Name `xml:"meta"`
	OsmBase string   `xml:"osm_base,attr"`
}

type Node struct {
	XMLName   xml.Name `xml:"node"`
	Id        int      `xml:"id,attr"`
	Latitude  float64  `xml:"lat,attr"`
	Longitude float64  `xml:"lon,attr"`
	Tags      []Tag    `xml:"tag"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func GetATMs() {
	//55.66076 / 37.480117 (lat/lon)
	//55.809556 / 37.7053288 (lat/lon)

	// Compose querry
	OSMserv := "https://overpass-api.de/api/interpreter?data="
	OSMdata := "node[amenity=atm](55.66076,37.480117,55.809556,37.7053288);out;"
	querry := OSMserv + OSMdata

	// Make request
	log.Println("Making request: ", querry)
	resp, err := http.Get(querry)
	check(err)
	defer resp.Body.Close()

	// Decode response body into struct
	log.Println("Decoding response into Osm struct...")
	var atms Osm
	xml.NewDecoder(resp.Body).Decode(&atms)

	// Write resulting struct into .xml file
	log.Println("Writing into osmOutput.xml...")
	xmlBody, err := xml.MarshalIndent(atms, "", "  ")
	check(err)
	err = os.WriteFile("osmOutput.xml", xmlBody, 0644)
	check(err)
}

func MakeQuerry() {
	// Make request
	querry := "https://overpass-api.de/api/interpreter?data=node(1);out;"
	log.Println("Making request: ", querry)
	resp, err := http.Get(querry)
	check(err)
	defer resp.Body.Close()

	// Write response to file
	log.Println("Writing into testResponseInterpreter.xml...")
	out, err := os.Create("testResponseInterpreter.xml")
	check(err)
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	check(err)
}
