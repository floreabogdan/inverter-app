package dashboard

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func (d *Dashboard) initRoutes() {
	d.core.WebEngine.LoadHTMLGlob("static/views/**/**")
	d.core.WebEngine.Use(static.Serve("/static/assets", static.LocalFile("static/assets", false)))

	d.core.WebEngine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{})
	})
}
