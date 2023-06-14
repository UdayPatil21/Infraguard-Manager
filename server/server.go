package main

import (
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
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
	r.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,origin,Cache-Control,X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
