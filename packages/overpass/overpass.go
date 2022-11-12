package overpass

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type City struct {
	Name     string
	Radius   float64
	Lat, Lon float64
}

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

func MakeQuerry(c City, amenity string) (Osm, error) {
	// Compose querry
	const OSMserv = "https://overpass-api.de/api/interpreter?data="
	OSMdata := fmt.Sprintf("node[amenity=\"%s\"](around:%f,%f,%f);out;", amenity, c.Radius, c.Lat, c.Lon)
	querry := OSMserv + OSMdata

	// Make querry
	// TODO: handle http errors
	resp, err := http.Get(querry)
	if err != nil {
		return Osm{}, err
	}
	defer resp.Body.Close()

	// Decode body into struct
	var osm Osm
	err = xml.NewDecoder(resp.Body).Decode(&osm)
	if err != nil {
		return Osm{}, err
	}

	return osm, nil
}
