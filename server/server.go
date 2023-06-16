package main

import (
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	"infraguard-manager/middleware"
	"infraguard-manager/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	// port := "4200"
	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	log.Println("Ping Called")
	// 	c.JSON(http.StatusOK, "Pong")
	// })

	//Add middleware CORS
	// r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware)

	//Init Config
	configHelper.InitConfig()

	//Init logger
	logger.Init()

	// Check all agent status concurrently
	// sheduler.CheckAgentStatus()

	//Initialize routes
	routes.InitRoutes(r)
	// log.Println("Server started on :", port)
	r.Run(":" + configHelper.GetString("Port"))

}
func main() {
	StartServer()
}
