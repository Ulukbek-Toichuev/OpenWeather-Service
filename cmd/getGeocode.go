package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	a "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
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

	var data []a.GeoCode

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
