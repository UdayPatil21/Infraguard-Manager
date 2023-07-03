package windows

import (
	"infraguard-manager/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWindowsRoutes(routeGroup *gin.RouterGroup) {

	r := routeGroup.Group("/platform/windows").Use(middleware.ValidateDomain())
	r.POST("/send-command", SendCommands)
}

func SendCommands(c *gin.Context) {
	log.Println("Welcome to windows")
	c.JSON(http.StatusOK, "done")
}
