package routers

import (
	"html/template"
	"net/http"
	"strings"

	apiJSON "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
	cmd "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
)

//Здесь будет код который рендерит HTML страницу
//Here will be the code that renders the HTML page

var (
	tmp         *template.Template
	owm         cmd.OpenWeather
	weatherOWM  apiJSON.CurrentWeather
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
		http.Redirect(w, r, "/", http.StatusBadRequest)
		/*data := Bind{
			Len: len(fCity),
		}
		tmp.ExecuteTemplate(w, "city.html", data)*/
	} else {

		city = fCity

		stringOut := string(owm.GetGeocode(city))

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

			s, weatherDesc := owm.GetWeatherStat(CurrentWeatherUrl)
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
}
