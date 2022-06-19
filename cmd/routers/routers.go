package routers

import (
	"net/http"

	currentData "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
	dbConnect "github.com/Ulukbek-Toychuev/OpenWeather-Service/db"
	"github.com/gin-gonic/gin"
)

func GinTest() {
	router := gin.Default()
	router.Static("/style", "./tmp/style")
	router.LoadHTMLGlob("tmp/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "main.html", nil)
	})

	router.GET("/City", func(ctx *gin.Context) {
		city := ctx.Request.FormValue("City")
		resData, stat := currentData.GetData(city)
		dbConnect.ConnectDB(city)
		aqiRes := currentData.GetAQI(&resData)
		if stat == 400 {
			ctx.HTML(http.StatusBadRequest, "cityBadReq.html", nil)
		} else {
			ctx.HTML(http.StatusOK, "city.html", gin.H{
				"Country":  resData.Location.Country,
				"City":     city,
				"CurrTemp": resData.Current.CurrentTemp,
				"Desc":     resData.Current.Condition.Text,
				"AQI":      aqiRes,
			})
		}

	})

	router.Run(":8000")

}
