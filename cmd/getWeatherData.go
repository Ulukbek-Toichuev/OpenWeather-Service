package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	a "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
)

const token string = "da303db859918e01a675709c157ca661"

//Интерфейс для методов которые взаимодействуют с Open Weather Map
//Interface for methods that interact with Open Weather Map
type GetOpenWeatherData interface {
	GetWeatherStat()
	GetAirPollution()
}

type OpenWeather struct {
}

//Функция для получения данных о погоде по указанному городу
//Function for getting weather data for the specified city
func (owm OpenWeather) GetWeatherStat(city string) float64 {
	lat, lon := getGeocode(city)

	CurrentWeatherUrl := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "6&appid=" + token

	requestWeather, err := http.Get(CurrentWeatherUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer requestWeather.Body.Close()

	respBodyWeather, err := ioutil.ReadAll(requestWeather.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherOWM a.CurrentWeather

	err = json.Unmarshal(respBodyWeather, &weatherOWM)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range weatherOWM.Weather {
		fmt.Println("-----------------")
		fmt.Println("Briefly about the weather:", p.WeatherMain)
		fmt.Println("Weather description:", p.WeatherDescription)
	}

	currentWeather := weatherOWM.Main.MainTempMax - 273.15

	fmt.Printf("Current temperature: %.2f\n", currentWeather)

	return currentWeather

}

//Функция для получения данных о загрязнении воздуха в указанном городе
//Function to get data about air pollution in the specified city
func (owm OpenWeather) GetAirPollution(city string) a.AirPollution {
	var lat, lon string
	lat, lon = getGeocode(city)

	currentAirPollutionURL := "http://api.openweathermap.org/data/2.5/air_pollution?lat=" + lat + "&lon=" + lon + "&appid=" + token + ""

	requestAir, err := http.Get(currentAirPollutionURL)

	if err != nil {
		log.Fatal(err)
	}
	defer requestAir.Body.Close()
	responseBodyAir, err := ioutil.ReadAll(requestAir.Body)

	if err != nil {
		log.Fatal(err)
	}

	var air a.AirPollution

	err = json.Unmarshal(responseBodyAir, &air)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(air.List)

	return air
}