package linux

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
	r := routeGroup.Group("/linux")
	r.POST("/send-command", sendCommand)
	r.POST("/execute-script:machineID", executeScript)
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
	machineID := c.Param("machineID")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("error reading request body", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	// if err != nil {
	// 	logger.Error("error binding data", err)
	// 	c.JSON(http.StatusExpectationFailed, err)
	// 	return
	// }
	input.Script = data
	input.MachineID = machineID
	res, err := executeScriptService(input)
	if err != nil {
		logger.Error("Error executing script", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	logger.Info("OUT:executeScript")
	c.JSON(http.StatusOK, res)
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
