package apiCall

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
)

type WeatherAPIResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func ApiCalltoOpenWeather(location string, apiKey string) {
	queryParams := url.Values{}

	queryParams.Add("q", location)
	queryParams.Add("APPID", apiKey)
	queryParams.Add("units", "imperial")
	url := "https://api.openweathermap.org/data/2.5/weather?" + queryParams.Encode()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response WeatherAPIResponse
	json.Unmarshal(resBody, &response)

	fmt.Println("Temperature: ", convertFahrenheitToCelsius(response.Main.Temp), "C, but feels like: ", convertFahrenheitToCelsius(response.Main.FeelsLike), "C")
}

func convertFahrenheitToCelsius(tempf float64) float64 {
	return math.Round((tempf - 32) / 1.8)
}
