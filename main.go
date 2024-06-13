package main

import (
	"log"
	"net/http"
	"weather-app/handler"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/weather/current", handler.GetCurrentWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(viper.Get("PORT").(string), r))
}
