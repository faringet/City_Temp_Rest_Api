package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Location struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
	Coord   struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
}

func ReturnMap() map[int]string {

	// Read the JSON file
	data, err := ioutil.ReadFile("city.list.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON data
	var locations []Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
	}

	// Create a map to store the values
	locationMap := make(map[int]string)
	for _, location := range locations {
		locationMap[location.ID] = location.Name
	}

	return locationMap
}

func GetCityID(m map[int]string, val string) int {
	for k, v := range m {
		if v == val {
			return k
		}
	}
	// Return an empty string if the value is not present
	return 0
}
