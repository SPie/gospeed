package speedtest

import "github.com/gin-gonic/gin"

// Controller handles Speedtest API calls
type Controller interface {
	GetSpeedtests() gin.HandlerFunc
}

type controller struct {
	repository Repository
}

// NewController creates a new Speedtest Controller
func NewController(repository Repository) Controller {
	return controller{repository: repository}
}

func (controller controller) GetSpeedtests() gin.HandlerFunc {
	return func(c *gin.Context) {
		speedtests := controller.repository.FindAll()

		c.JSON(200, map[string][]Speedtest{"speedtests": speedtests})
	}
}
