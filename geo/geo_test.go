package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetCurrentLocation(t *testing.T) {
	// Arrange block
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}

	// Act block
	got, err := geo.GetCurrentLocation(city)

	// Assert block
	// Проверяем ошибку получения города
	if err != nil {
		t.Error(err.Error())
	}

	// Проверяем, что полученный результат соответствует ожидаемому
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получили %v", expected, got)
	}
}
