package main

import (
	s "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
)

func main() {

	var owm s.OpenWeather

	owm.GetAirPollution()
	/*router := gin.Default()
	router.LoadHTMLGlob("tmp/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")*/

}
