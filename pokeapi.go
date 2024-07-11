package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
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

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreaResponse LocationAreaResponse

	err = json.Unmarshal(data, &locationAreaResponse)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	// fmt.Println(locationAreaResponse.Results[4].Name)
	return locationAreaResponse, nil

}
