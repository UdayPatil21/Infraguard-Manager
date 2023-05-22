package linux

import (
	"bytes"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io"
	"net/http"

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
	}
	out, err := sendCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Info("OUT:sendCommand")
	c.JSON(http.StatusOK, out)
}

func executeScript(c *gin.Context) {
	logger.Info("IN:executeScript")

	//Read form data
	per := c.Request.FormValue("permission")
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("Error getting form data", err)
		c.JSON(http.StatusBadRequest, err)
	}
	content, _ := file.Open()
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, content); err != nil {
		c.JSON(http.StatusExpectationFailed, err)
	}
	var input model.Executable
	input.Permission = per
	input.Script = buf.Bytes()

	res, err := executeScriptService(input)
	if err != nil {
		logger.Error("Error executing script file", err)
		c.JSON(http.StatusBadRequest, err)
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
	}
	out, err := sudoCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	logger.Info("OUT:sudoCommand")
	c.JSON(http.StatusOK, out)
}
