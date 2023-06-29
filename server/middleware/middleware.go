package middleware

import (
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	"net/http"

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

func Auth(tokenString string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if tokenString == "" {
			logger.Error("Token string not Generated")
			context.Abort()
			return
		}
		context.Request.Header.Add("Authorization", tokenString)
		context.Next()
	}
}
