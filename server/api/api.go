package api

import (
	model "infraguard-manager/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterInstance(c *gin.Context) {
	//Register all new instances
	var instanceInfo model.InstanceInfo

	err := c.Bind(&instanceInfo)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	c.JSON(http.StatusOK, "success")

}
