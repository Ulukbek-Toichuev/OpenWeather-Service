package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jsonSchema "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
)

const token string = "da303db859918e01a675709c157ca661"

//Интерфейс для методов которые взаимодействуют с Open Weather Map
//Interface for methods that interact with Open Weather Map
type GetOpenWeatherData interface {
	GetWeatherStat()
	GetAirPollution()
	GetGeocode()
}

type OpenWeather struct {
}

//Функция для получения названия города по геокоду.
//Function for getting cities name for specified geocode.

func (owm OpenWeather) GetGeocode(city string) []byte {
	var data []jsonSchema.GeoCode
	geoCodeUrl := "http://api.openweathermap.org/geo/1.0/direct?q=" + city + ",&appid="
	geoCodeUrl = geoCodeUrl + token

	requestGeocode, err := http.Get(geoCodeUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(requestGeocode.Request.Method, geoCodeUrl, requestGeocode.Status)
	defer requestGeocode.Body.Close()

	geocodeBody, err := ioutil.ReadAll(requestGeocode.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(geocodeBody, &data)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

//Функция для получения данных о погоде по указанному городу.
//Function for getting weather data for the specified city.
func (owm OpenWeather) GetWeatherStat(city string) (string, string) {
	var weatherOWM jsonSchema.CurrentWeather
	var weatherDesc string

	requestWeather, err := http.Get(city)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(requestWeather.Request.Method, city, requestWeather.Status)
	defer requestWeather.Body.Close()

	respBodyWeather, err := ioutil.ReadAll(requestWeather.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(respBodyWeather, &weatherOWM)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range weatherOWM.Weather {
		weatherDesc = p.WeatherDescription
	}

	s := fmt.Sprintf("%.2f", weatherOWM.Main.MainTemp-273.15)

	return s, weatherDesc
}

//Функция для получения данных о загрязнении воздуха в указанном городе.
//Function to get data about air pollution in the specified city.
func (owm OpenWeather) GetAirPollution(city string) jsonSchema.AirPollution {

	requestAir, err := http.Get(city)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(requestAir.Request.Method, city, requestAir.Status)
	defer requestAir.Body.Close()

	responseBodyAir, err := ioutil.ReadAll(requestAir.Body)
	if err != nil {
		log.Fatal(err)
	}

	var air jsonSchema.AirPollution

	err = json.Unmarshal(responseBodyAir, &air)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(air.List)

	return air
}
