package geo_test

import (
	"testing"
	"weather/geo"
)

// Позитивный тест
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

// Негативный тест
func TestGetCurrentLocationNoCity(t *testing.T) {
	city := "Londop"
	_, err := geo.GetCurrentLocation(city)

	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получили %v", geo.ErrNoCity, err)
	}

	if err != geo.ErrNotOk {
		t.Errorf("Ожидалось %v, получили %v", geo.ErrNotOk, err)
	}
}
