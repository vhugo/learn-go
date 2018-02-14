// Facade
//
// It shields the code from unwanted access, orders some calls, and hides the
// complexity scope from the user.
//
// You use Facade when you want to hide the complexity of some tasks, especially
// when most of them share utilities (such as authentication in an API). A
// library is a form of facade, where someone has to provide some methods for a
// developer to do certain things in a friendly way. This way, if a developer
// needs to use your library, he doesn't need to know all the inner tasks to
// retrieve the result he/she wants.
//
// So, you use the Facade design pattern in the following scenarios:
//
// - When you want to decrease the complexity of some parts of our code. You hide
// that complexity behind the facade by providing a more easy-to-use method.
//
// - When you want to group actions that are cross-related in a single place.
//
// - When you want to build a library so that others can use your products without
// worrying about how it all works.
//
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetBeyGeoCoordinates(lat, lon float32) (Weather, error)
}

type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}

type CurrentWeatherData struct {
	APIkey string
}

func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(let, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s",
			city, countryCode, c.APIkey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s",
			lat, lon, c.APIkey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutils.ReadAll(resp.body)
		if errMsg != nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}

		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n",
			resp.StatusCode, errMsg)

		return
	}

	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()

	return
}
