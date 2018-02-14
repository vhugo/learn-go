// Acceptance criteria
//
// The OpenWeatherMap API gives lots of information, so we are going to focus on
// getting live weather data in one city in some geo-located place by using its
// latitude and longitude values. The following are the requirements and
// acceptance criteria for this design pattern:
//
// 1. Provide a single type to access the data. All information retrieved from
// OpenWeatherMap service will pass through it.
//
// 2. Create a way to get the weather data for some city of some country.
//
// 3. Create a way to get the weather data for some latitude and longitude
// position.
//
// 4. Only second and thrird point must be visible outside of the package;
// everything else must be hidden (including all connection-related data).
//
package main

import (
	"bytes"
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{
		"coord": {
				"lon": -3.7,
				"lat": 40.42
		},
		"weather": [{
				"id": 803,
				"main": "Clouds",
				"description": "broken clouds",
				"icon": "04n"
		}],
		"base": "stations",
		"main": {
				"temp": 303.56,
				"pressure": 1016.46,
				"humidity": 26.8,
				"temp_min": 300.95,
				"temp_max": 305.93
		},
		"wind": {
				"speed": 3.17,
				"deg": 151.001
		},
		"rain": {
				"3h": 0.0075
		},
		"clouds": {
				"all": 68
		},
		"dt": 1471295823,
		"sys": {
				"type": 3,
				"id": 1442829648,
				"message": 0.0278,
				"country": "ES",
				"sunrise": 1471238808,
				"sunset": 1471288232
		},
		"id": 3117735,
		"name": "Madrid",
		"cod": 200
}`

	r := bytes.NewReader([]byte(response))
	return r
}

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}
