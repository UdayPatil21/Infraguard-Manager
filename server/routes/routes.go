package routes

import (
	"infraguard-manager/api"
	"infraguard-manager/api/linux"
	"infraguard-manager/api/windows"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	//Initialize all required routes

	routeGroup := route.Group("/api")
	routeGroup.POST("/instance-info", api.RegisterInstance)
	windows.InitWindowsRoutes(routeGroup)
	linux.InitLinuxRoutes(routeGroup)
}
