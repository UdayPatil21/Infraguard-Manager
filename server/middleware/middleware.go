package middleware

import (
	"infraguard-manager/helpers/configHelper"
	model "infraguard-manager/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	// c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", configHelper.GetString("Infraguard-URL"))
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "*")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}

}

func ValidateDomain() gin.HandlerFunc {
	return func(context *gin.Context) {
		response := model.Response{}
		//Statndard Output
		response.Status = false
		hostString := configHelper.GetString("Infraguard-URL")
		// Check for authorized domain
		if !strings.Contains(context.Request.Host, hostString) && !strings.Contains(context.Request.Host, "3.0.119.137") && !strings.Contains(context.Request.Host, "localhost") {
			response.Data = "Unknown Server"
			response.Error = "Unknown Server Connecting .... Error"
			context.JSON(http.StatusBadRequest, response)
			context.Abort()
			return
		}
		context.Next()
	}
}
