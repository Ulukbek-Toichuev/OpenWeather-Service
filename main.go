package main

import (
	//s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
	s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd/routers"
)

func main() {
	s.Server()
	//var owm s.OpenWeather

	//owm.GetWeatherStat()

	//owm.GetWeatherStat()
	/*router := gin.Default()
	router.LoadHTMLGlob("tmp/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")*/

}
