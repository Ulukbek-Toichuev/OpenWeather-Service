package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const token string = "da303db859918e01a675709c157ca661"

var (
	url  string = "http://api.openweathermap.org/geo/1.0/direct?q=,&appid="
	city string
)

type Hueta struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func buildURL() string {
	fmt.Print("Enter your city: ")
	fmt.Scanln(&city)
	url = "http://api.openweathermap.org/geo/1.0/direct?q=" + city + ",&appid="
	url = url + token

	return url
}

func GetGeocode() {

	buildURL()
	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Hueta

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
