package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const token string = "da303db859918e01a675709c157ca661"

var (
	geoCodeURL string = "http://api.openweathermap.org/geo/1.0/direct?q=,&appid="
	cityGeoi   string
)

func buildURL() string {
	fmt.Print("Enter your city: ")
	fmt.Scanln(&cityGeoi)
	geoCodeURL = "http://api.openweathermap.org/geo/1.0/direct?q=" + cityGeoi + ",&appid="
	geoCodeURL = geoCodeURL + token

	return geoCodeURL
}

func GetGeocode() {

	buildURL()
	fmt.Println(geoCodeURL)

	resp, err := http.Get(geoCodeURL)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
