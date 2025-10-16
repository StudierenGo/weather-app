package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const getCurrentURL = "https://ipapi.co/json/"
const checkCityUrl = "https://countriesnow.space/apu/v0.1/countries/population/"

type GeoData struct {
	City string `json:"city"`
}

type CityCheck struct {
	Error bool `json:"error"`
}

func GetCurrentLocation(city string) (*GeoData, error) {
	var geoData GeoData

	if city != "" {
		isCity := checkCity(city)

		if !isCity {
			return nil, errors.New("NOCITY")
		}

		return &GeoData{
			City: city,
		}, nil
	}

	response, err := http.Get(getCurrentURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &geoData)

	return &geoData, nil
}

func checkCity(city string) bool {
	var checkCityResponse CityCheck

	body, err := json.Marshal(map[string]string{
		"city": city,
	})
	if err != nil {
		return false
	}

	response, err := http.Post(checkCityUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	json.Unmarshal(body, &checkCityResponse)
	return checkCityResponse.Error
}
