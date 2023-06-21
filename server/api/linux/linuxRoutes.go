package linux

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
	r := routeGroup.Group("/linux")
	r.POST("/send-command", sendCommand)
	r.POST("/execute-script", executeScript)
	r.POST("/sudo-command", sudoCommand)
}

func sendCommand(c *gin.Context) {
	logger.Info("IN:sendCommand")
	input := model.RunCommand{}
	err := c.Bind(&input)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	out, err := sendCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}
	logger.Info("OUT:sendCommand")
	c.JSON(http.StatusOK, out)
}

func executeScript(c *gin.Context) {
	logger.Info("IN:executeScript")

	var input model.Executable
	err := c.Bind(&input)
	if err != nil {
		logger.Error("error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	// res, err := executeScriptService(input)
	res, err := executeScriptService(input)
	if err != nil {
		logger.Error("Error executing script", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// s := SanitizeScript(res)
	logger.Info("OUT:executeScript")
	c.JSON(http.StatusOK, res)
}
func SanitizeScript(script string) string {
	s2 := strings.Replace(script, `\n`, "\n", -1)
	// s := strings.ReplaceAll(s2, "\\", "")
	return s2
}
func sudoCommand(c *gin.Context) {
	logger.Info("IN:sudoCommand")
	input := model.RunCommand{}
	err := c.Bind(&input)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	out, err := sudoCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:sudoCommand")
	c.JSON(http.StatusOK, out)
}
