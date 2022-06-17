package routers

import (
	"net/http"

	currentData "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
	"github.com/gin-gonic/gin"
)

/*
func FiberTest() {
	engine := html.New("./tmp", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/style", "./tmp/style")
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("main", fiber.Map{
			"City": "City",
		})
	})

	app.Get("/City", func(c *fiber.Ctx) error {
		city := c.FormValue("City")
		resData, stat := currentData.GetData(city)
		aqiRes := currentData.GetAQI(&resData)
		if stat == 400 {
			return c.Render("cityBadReq", c.SendStatus(400))
		}
		return c.Render("city", fiber.Map{
			"Country":  resData.Location.Country,
			"City":     city,
			"CurrTemp": resData.Current.CurrentTemp,
			"Desc":     resData.Current.Condition.Text,
			"AQI":      aqiRes,
		})
	})

	log.Fatal(app.Listen(":80"))
}*/

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
