package routes

import (
	"infraguard-manager/api"
	activation "infraguard-manager/api/agent-activation"
	"infraguard-manager/api/linux"
	"infraguard-manager/api/windows"
	helper "infraguard-manager/helpers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	//Initialize all required routes

	routeGroup := route.Group("/api/agent")
	routeGroup.POST("/registration/serverinfo", api.RegisterInstance)
	routeGroup.POST("/update/serverinfo", api.UpdateServerInfo)
	routeGroup.POST("/verify/connectionstatus", helper.CheckStatus)
	routeGroup.POST("/addAgentActivation", activation.AddAgentActivation)
	routeGroup.GET("/getAgentActivation:id", activation.GetAgentActivationById)
	routeGroup.GET("/getAllActivations", activation.GetAllActivation)
	routeGroup.POST("/editAgentActivation", activation.UpdateAgentActivation)
	routeGroup.GET("/deleteAgentActivation:id", activation.DeleteAgentActivationById)
	routeGroup.GET("/getAllServers", activation.GetAllServers)
	windows.InitWindowsRoutes(routeGroup)
	linux.InitLinuxRoutes(routeGroup)
}
