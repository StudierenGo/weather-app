package main

import (
	"flag"
	"fmt"
	"weather/geo"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	// format := flag.Int("format", 1, "Формат вывода")

	flag.Parse()

	geoData, err := geo.GetCurrentLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(*geoData)
}
