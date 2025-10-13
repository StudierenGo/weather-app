package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"weather/geo"
)

const BASE_URL = "https://wttr.in/"

func GetCurrentWeather(geoData geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse(BASE_URL + geoData.City)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("format", strconv.Itoa(format))

	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(body), nil
}
