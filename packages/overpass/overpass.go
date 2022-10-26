package overpass

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
