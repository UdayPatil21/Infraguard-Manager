package linux

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
	r := routeGroup.Group("/linux")
	r.POST("/send-command", SendCommands)
}

func SendCommands(c *gin.Context) {
	log.Println("Welcome to linux")
	c.JSON(http.StatusOK, "done")
}
