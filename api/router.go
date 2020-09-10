package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spie/gospeed/speedtest"
)

// Run starts the router engine
func Run(host string, port string, speedtestController speedtest.Controller) {
	engine := gin.Default()

	api := engine.Group("/api")
	{
		api.GET("/speedtests", speedtestController.GetSpeedtests())
	}

	engine.Run(fmt.Sprintf("%s:%s", host, port))
}
