package routes

import (
	"infraguard-manager/api"
	activation "infraguard-manager/api/agent-activation"
	"infraguard-manager/api/linux"
	"infraguard-manager/api/windows"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	//Initialize all required routes

	routeGroup := route.Group("/api")
	routeGroup.POST("/instance-info", api.RegisterInstance)
	routeGroup.POST("/update-ip", api.UpdateAgent)
	routeGroup.POST("/addAgentActivation", activation.AddAgentActivation)
	routeGroup.GET("/getAgentActivation:id", activation.GetAgentActivationById)
	routeGroup.GET("/getAllActivations", activation.GetAllActivation)
	routeGroup.POST("/editAgentActivation", activation.UpdateAgentActivation)
	routeGroup.GET("/deleteAgentActivation:id", activation.DeleteAgentActivationById)
	windows.InitWindowsRoutes(routeGroup)
	linux.InitLinuxRoutes(routeGroup)
}
