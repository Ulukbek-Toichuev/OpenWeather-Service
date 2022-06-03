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

func GetGeocode() (string, string) {

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

	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	stringOut := string(out)

	res := strings.Split(stringOut, ",")

	var lat, lon string = res[0], res[1]

	lat = strings.ReplaceAll(lat, "[{\"lat\":", "")
	lon = strings.ReplaceAll(lon, "\"lon\":", "")
	lon = strings.ReplaceAll(lon, "}]", "")

	return lat, lon
}
