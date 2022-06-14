package routers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	a "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
	s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
)

//Здесь будет код который рендерит HTML страницу
//Here will be the code that renders the HTML page

var (
	tmp         *template.Template
	owm         s.OpenWeather
	weatherOWM  a.CurrentWeather
	lat         string
	lon         string
	city        string
	weatherDesc string
)

type Bind struct {
	Len int
}

const token string = "da303db859918e01a675709c157ca661"

func init() {
	tmp = template.Must(template.ParseGlob("tmp/*.html"))
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("tmp/style"))))
}

func Server() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/City", SelectCity)
	http.ListenAndServe(":80", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "main.html", nil)
}

func SelectCity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fCity := r.FormValue("City")
	if len(fCity) < 2 {
		data := Bind{
			Len: len(fCity),
		}
		tmp.ExecuteTemplate(w, "city.html", data)
	} else {

		city = fCity

		GeoCodeUrl := "http://api.openweathermap.org/geo/1.0/direct?q=" + city + ",&appid="
		GeoCodeUrl = GeoCodeUrl + token
		fmt.Println(token)

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
		if len(res) < 2 {
			data := Bind{
				Len: len(res),
			}
			tmp.ExecuteTemplate(w, "city.html", data)
		} else {
			lat, lon = res[0], res[1]
			lat = strings.ReplaceAll(lat, "[{\"lat\":", "")
			lon = strings.ReplaceAll(lon, "\"lon\":", "")
			lon = strings.ReplaceAll(lon, "}]", "")

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

			err = json.Unmarshal(respBodyWeather, &weatherOWM)
			if err != nil {
				log.Fatal(err)
			}

			currentWeather := weatherOWM.Main.MainTempMax - 273.15

			for _, p := range weatherOWM.Weather {
				weatherDesc = p.WeatherDescription
			}

			s := fmt.Sprintf("%.2f", currentWeather)
			Len := len(res)

			c := struct {
				TempCur  string
				City     string
				Describe string
				Len      int
			}{
				TempCur:  s,
				City:     fCity,
				Describe: weatherDesc,
				Len:      Len,
			}

			tmp.ExecuteTemplate(w, "city.html", c)
		}

	}
	//temp, weathDesc, l := owm.GetWeatherStat(fCity)
	// Begin

}
