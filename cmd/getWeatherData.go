package cmd

const token string = "da303db859918e01a675709c157ca661"

//Интерфейс для методов которые взаимодействуют с Open Weather Map
//Interface for methods that interact with Open Weather Map
type GetOpenWeatherData interface {
	GetWeatherStat()
	GetAirPollution()
}

type OpenWeather struct {
}

//Функция для получения данных о погоде по указанному городу.
//Function for getting weather data for the specified city.
/*func (owm OpenWeather) GetWeatherStat(city string) (float64, string, int) {
	lat, lon, l := getGeocode(city)
	fmt.Println("*** FUNC GetWeatherStat ***", l)

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
	var weatherDesc string

	err = json.Unmarshal(respBodyWeather, &weatherOWM)
	if err != nil {
		log.Fatal(err)
	}

	currentWeather := weatherOWM.Main.MainTempMax - 273.15

	for _, p := range weatherOWM.Weather {
		weatherDesc = p.WeatherDescription
	}

	fmt.Printf("Current temperature: %.2f\n", currentWeather)

	return currentWeather, weatherDesc, l

}

//Функция для получения данных о загрязнении воздуха в указанном городе.
//Function to get data about air pollution in the specified city.
func (owm OpenWeather) GetAirPollution(city string) (a.AirPollution, int) {
	var lat, lon string
	lat, lon, l := getGeocode(city)

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

	return air, l
}*/
