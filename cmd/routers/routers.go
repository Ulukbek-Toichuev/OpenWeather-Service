package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Здесь будет код который рендерит HTML страницу
//Here will be the code that renders the HTML page
func Router() {
	router := gin.Default()
	router.LoadHTMLGlob("tmp/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8080")
}
