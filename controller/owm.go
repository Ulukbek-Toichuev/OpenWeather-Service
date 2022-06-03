package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const token string = "da303db859918e01a675709c157ca661"

var (
	GeoCodeUrl        string = "http://api.openweathermap.org/geo/1.0/direct?q=,&appid="
	CurrentWeatherUrl string = "https://api.openweathermap.org/data/2.5/weather?lat=&lon=&appid="
	city              string
)

func getGeocode() (string, string) {
	var lat, lon string

	fmt.Print("Enter your city: ")
	fmt.Scanln(&city)
	GeoCodeUrl = "http://api.openweathermap.org/geo/1.0/direct?q=" + city + ",&appid="
	GeoCodeUrl = GeoCodeUrl + token

	resp, err := http.Get(GeoCodeUrl)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []GeoCode

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	stringOut := string(out)

	res := strings.Split(stringOut, ",")

	lat, lon = res[0], res[1]

	lat = strings.ReplaceAll(lat, "[{\"lat\":", "")
	lon = strings.ReplaceAll(lon, "\"lon\":", "")
	lon = strings.ReplaceAll(lon, "}]", "")

	return lat, lon
}

func GetWeatherStat() {
	var lat, lon string
	lat, lon = getGeocode()

	CurrentWeatherUrl = "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "6&appid=" + token

	requestOWM, err := http.Get(CurrentWeatherUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer requestOWM.Body.Close()

	respBodyOWM, err := ioutil.ReadAll(requestOWM.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherOWM OWM

	err = json.Unmarshal(respBodyOWM, &weatherOWM)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range weatherOWM.Weather {
		fmt.Println("-----------------")
		fmt.Println("Briefly about the weather:", p.WeatherMain)
		fmt.Println("Weather description:", p.WeatherDescription)
	}

	fmt.Printf("Current temperature: %.2f\n", weatherOWM.Main.MainTempMax-273.15)

}
