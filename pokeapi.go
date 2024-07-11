package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ShahSau/pokedexcli/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	fullUrl := baseUrl + "/location-area"

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	dat, ok := c.cache.Get(fullUrl)

	if ok {

		fmt.Println("Cache hit")
		var locationAreaResponse LocationAreaResponse

		err := json.Unmarshal(dat, &locationAreaResponse)

		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreaResponse, nil

	}

	fmt.Println("Cache miss")
	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreaResponse LocationAreaResponse

	err = json.Unmarshal(dat, &locationAreaResponse)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullUrl, dat)

	// fmt.Println(locationAreaResponse.Results[4].Name)
	return locationAreaResponse, nil

}
