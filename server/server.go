package main

import (
	"errors"
	"infraguard-manager/db"
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	logs "infraguard-manager/helpers/logger"
	"infraguard-manager/middleware"
	"infraguard-manager/middleware/auth"
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

	//Init Config
	configHelper.InitConfig()

	//Init logger
	// logs.Logger.Init()
	logs.Initialize()

	err := errors.New("erororororororo")
	logger.Logger.Sugar().Errorf("Test logs", err)
	//Generate JWT Tokens
	auth.Init()

	//Add middleware CORS
	// r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware)

	//Login to infraguard.io and initiate the Authentication process
	//Then generate the user token
	//Used that user token to access all restricted api's of infraguard server
	auth.InfraLogin()

	//Initialize routes
	routes.InitRoutes(r)

	//Initialized mysql connection
	db.Init()

	// Check all agent status concurrently
	// sheduler.Scheduler()

	// log.Println("Server started on :", port)
	r.Run(":" + configHelper.GetString("Port"))

}
func main() {
	StartServer()
}
