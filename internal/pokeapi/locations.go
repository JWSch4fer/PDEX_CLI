package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	// this is to use next and previous
	if pageURL != nil {
		fullURL = *pageURL
	}

	//check if data is in the cache
	data, ok := c.cache.Get(fullURL)
	if !ok {
		fmt.Println("Cache miss -_-")

		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		resp, err := c.httpclient.Do(req)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return LocationAreasResponse{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		c.cache.Add(fullURL, data)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		locationAreasResp := LocationAreasResponse{}
		err = json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResp, nil

	} else {
		//data was stored in the cache!!!
		fmt.Println("Cache hit ^_^")
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResp, nil

	}
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

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check if data is in the cache
	data, ok := c.cache.Get(fullURL)
	if !ok {
		fmt.Println("Cache miss -_-")

		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpclient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return LocationArea{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		c.cache.Add(fullURL, data)
		if err != nil {
			return LocationArea{}, err
		}
		locationArea := LocationArea{}
		err = json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil

	} else {
		//data was stored in the cache!!!
		fmt.Println("Cache hit ^_^")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil

	}
}
