package geo

import (
	"encoding/json"
	"io"
	"net/http"
)

const BASE_URL = "https://ipapi.co/json/"

type GeoData struct {
	City string `json:"city"`
}

func GetCurrentLocation(city string) (*GeoData, error) {
	var geoData GeoData

	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}

	response, err := http.Get(BASE_URL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &geoData)

	return &geoData, nil
}
