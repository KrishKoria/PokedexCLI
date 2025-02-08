package api

import (
	"encoding/json"
	"net/http"
	"time"

	cache "github.com/KrishKoria/PokeCache"
)
var locationCache = cache.NewCache(5 * time.Minute)

type LocationResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Next string `json:"next"`
	Previous string `json:"previous"`
}

func FetchLocations(url string) (*LocationResponse, error) {
	if data, ok := locationCache.Get(url); ok {
        var locationResponse LocationResponse
        if err := json.Unmarshal(data, &locationResponse); err != nil {
            return nil, err
        }
        return &locationResponse, nil
    }

	res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var locationResponse LocationResponse
    if err := json.NewDecoder(res.Body).Decode(&locationResponse); err != nil {
        return nil, err
    }
	data, err := json.Marshal(locationResponse)
	if err != nil {
		return nil, err
	}
	locationCache.Add(url, data)

    return &locationResponse, nil
}
		