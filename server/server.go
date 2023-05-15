package main

import (
	"infraguard-manager/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	port := "4200"
	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	log.Println("Ping Called")
	// 	c.JSON(http.StatusOK, "Pong")
	// })

	//Initialize routes
	routes.InitRoutes(r)
	// log.Println("Server started on :", port)
	r.Run(":" + port)

}
func main() {

	StartServer()
}
