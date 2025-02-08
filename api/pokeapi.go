package api

import (
	"encoding/json"
	"net/http"
)

type LocationResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Next string `json:"next"`
	Previous string `json:"previous"`
}

func FetchLocations(url string) (*LocationResponse, error) {
	res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var locationResponse LocationResponse
    if err := json.NewDecoder(res.Body).Decode(&locationResponse); err != nil {
        return nil, err
    }

    return &locationResponse, nil
}
		