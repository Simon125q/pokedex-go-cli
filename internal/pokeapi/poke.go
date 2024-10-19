package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL string = "https://pokeapi.co/api/v2/location-area"

type AreaResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getLocationAreas() []Location {
	res, err := http.Get(baseURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Println(res.Status)
		return nil
	}
	decoder := json.NewDecoder(res.Body)
	var items []Location
	if err := decoder.Decode(&items); err != nil {
		return nil
	}
	return items
}
