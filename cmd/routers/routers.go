package routers

import (
	"fmt"
	"html/template"
	"net/http"

	a "github.com/Ulukbek-Toychuev/OpenWeather-Service/api"
	s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
)

//Здесь будет код который рендерит HTML страницу
//Here will be the code that renders the HTML page

var tmp *template.Template
var owm s.OpenWeather
var weatherOWM a.CurrentWeather

func init() {
	tmp = template.Must(template.ParseGlob("tmp/*.html"))
}

func Server() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/City", selectCity)
	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "main.html", nil)
}

func selectCity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fCity := r.FormValue("City")

	temp := owm.GetWeatherStat(fCity)
	//owm.GetWeatherStat(fCity)

	s := fmt.Sprintf("%f", temp)

	c := struct {
		TempCur string
	}{
		TempCur: s,
	}

	tmp.ExecuteTemplate(w, "city.html", c)
}
