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

var (
	tmp        *template.Template
	owm        s.OpenWeather
	weatherOWM a.CurrentWeather
)

func init() {
	tmp = template.Must(template.ParseGlob("tmp/*.html"))
	http.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("tmp/style"))))
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

	temp, weathDesc := owm.GetWeatherStat(fCity)

	s := fmt.Sprintf("%.2f", temp)

	c := struct {
		TempCur  string
		City     string
		Describe string
	}{
		TempCur:  s,
		City:     fCity,
		Describe: weathDesc,
	}

	tmp.ExecuteTemplate(w, "city.html", c)
}
