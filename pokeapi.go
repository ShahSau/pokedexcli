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
	Count    int                   `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []LocationAreaDetails `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationAreaDetails struct {
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

	return locationAreaResponse, nil

}

func (c *Client) GetDetailsOfLocationArea(locationAreaName string) (LocationArea, error) {
	fullUrl := baseUrl + "/location-area/" + locationAreaName

	dat, ok := c.cache.Get(fullUrl)

	if ok {
		fmt.Println("Cache hit")
		var locationArea LocationArea

		err := json.Unmarshal(dat, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil

	}

	fmt.Println("Cache miss")
	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationArea{}, err
	}

	var locationArea LocationArea

	err = json.Unmarshal(dat, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, dat)

	return locationArea, nil

}
