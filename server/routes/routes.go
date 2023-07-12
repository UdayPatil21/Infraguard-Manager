package routes

import (
	"infraguard-manager/api"
	"infraguard-manager/api/agent"
	activation "infraguard-manager/api/agent-activation"
	helper "infraguard-manager/helpers"
	"infraguard-manager/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.Engine) {
	//Initialize all required routes

	routeGroup := route.Group("/api/agent").Use(middleware.ValidateDomain())
	routeGroup.POST("/registration/serverinfo", api.RegisterInstance)
	routeGroup.POST("/update/serverinfo", api.UpdateServerInfo)
	routeGroup.POST("/verify/connectionstatus", helper.CheckStatus)
	routeGroup.POST("/addAgentActivation", activation.AddAgentActivation)
	routeGroup.GET("/getAgentActivation:id", activation.GetAgentActivationById)
	routeGroup.GET("/getAllActivations", activation.GetAllActivation)
	routeGroup.POST("/editAgentActivation", activation.UpdateAgentActivation)
	routeGroup.GET("/deleteAgentActivation:id", activation.DeleteAgentActivationById)
	routeGroup.GET("/getAllServers", activation.GetAllServers)
	// Create seperate routes for different OS support
	// r := route.Group("/api/agent")
	routeGroup.POST("/platform/linux/script/execute", agent.ExecuteScript)

	// windows.InitWindowsRoutes(r)
	// linux.InitLinuxRoutes(r)
}
