package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"weather-app/pkg/weather"

	"github.com/spf13/viper"
)

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	client := weather.NewClient(viper.Get("WEATHER_URL").(string), viper.Get("WEATHER_KEY").(string), time.Second*10)
	lat, long := r.URL.Query().Get("lat"), r.URL.Query().Get("long")
	weatherResponse, err := client.GetCurrentWeatherResponse(lat, long)
	if err != nil || weatherResponse == nil || len(weatherResponse.Weather) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := make(map[string]string)
	response["Current condition"] = weatherResponse.Weather[0].Main
	response["Description"] = weatherResponse.Weather[0].Description
	response["Current Temperature"] = fmt.Sprintf("%f", weatherResponse.Main.Temp)
	if weatherResponse.Rain != nil {
		response["Rain in the last 1 hour"] = weatherResponse.Rain.H1
	}
	if weatherResponse.Snow != nil {
		response["Snow in the last 1 hour"] = weatherResponse.Snow.H1
	}
	if weatherResponse != nil {
		json.NewEncoder(w).Encode(response)
	}
}
