package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Credentials", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,origin,Cache-Control,X-Requested-With")
	// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "*")

	// Second, we handle the OPTIONS problem
	// if c.Request.Method != "OPTIONS" {
	// 	c.Next()
	// } else {
	// 	// Everytime we receive an OPTIONS request,
	// 	// we just return an HTTP 200 Status Code
	// 	// Like this, Angular can now do the real
	// 	// request using any other method than OPTIONS
	// 	c.AbortWithStatus(http.StatusOK)
	// }

}
