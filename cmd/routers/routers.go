package routers

import (
	"log"

	currentData "github.com/Ulukbek-Toychuev/OpenWeather-Service/cmd"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
)

func FiberTest() {
	engine := html.New("./tmp", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
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
}
