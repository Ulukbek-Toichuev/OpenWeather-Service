package routers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	currentData "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
)

//Здесь будет код который рендерит HTML страницу
//Here will be the code that renders the HTML page

var tmp *template.Template

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
	city := r.FormValue("City")
	if len(city) < 2 {
		fmt.Println("ERROR")
	}
	resData := currentData.GetData(city)
	aqiRes := currentData.GetAQI(&resData)
	log.Println(aqiRes)
	c := struct {
		CurrTemp float32
		City     string
		Country  string
		Desc     string
		AQI      string
	}{
		CurrTemp: resData.Current.CurrentTemp,
		City:     city,
		Country:  resData.Location.Country,
		Desc:     resData.Current.Condition.Text,
		AQI:      aqiRes,
	}
	tmp.ExecuteTemplate(w, "city.html", c)

}
