package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	jsonSchema "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
)

const (
	token string = "81f299fabd474c9a80c114658221606"
	url   string = "http://api.weatherapi.com/v1/current.json?key=" + token + "&q="
	aqi   string = "&aqi=yes"
)

func GetData(city string) (jsonSchema.Data, int) {
	var currData jsonSchema.Data

	//Создаем URL
	getCityData := url + city + aqi

	//Отправляем запрос
	request, err := http.Get(getCityData)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(request.Request.Method, request.Status)
	defer request.Body.Close()

	//Читаем и записываем данные с тела
	respBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	//Сохраняем данные в структуру
	err = json.Unmarshal(respBody, &currData)
	if err != nil {
		log.Fatal(err)
	}
	return currData, request.StatusCode
}
