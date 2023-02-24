package api

import (
	"github.com/gin-gonic/gin"
)

func (d *Api) initRoutes() {
	d.core.WebEngine.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	d.core.WebEngine.GET("/api/inverter", func(c *gin.Context) {
		var inverter, ok = d.core.Inverter["dummyInverter"]
		if !ok {
			log.Error("Could not load inverter!")

			c.JSON(200, gin.H{"response": "Inverter not found!"})

			return
		}

		c.JSON(200, gin.H{
			"response": inverter.GetData(),
			"config":   inverter.GetConfig(),
		})
	})
}
