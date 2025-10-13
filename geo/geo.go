package geo

import (
	"encoding/json"
	"io"
	"net/http"
)

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

	response, err := http.Get("https://ipapi.co/json/")
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
