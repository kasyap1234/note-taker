package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Photo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
	Src Src    `json:"src"`
}

type Src struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type PexelResponse struct {
	Photos []Photo `json:"photos"`
}

// query == title default ;
func GetPhotos(query string) ([]Photo, error) {
	apiKey := os.Getenv("PEXELS_API_KEY")
	if apiKey == "" {
		return nil, errors.New("PEXELS_API_KEY environment variable not set")
	}
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&per_page=1", query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", apiKey)
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call the api: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pexels api returned status: %s", resp.Status)
	}

	var res PexelResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	return res.Photos, nil
}
