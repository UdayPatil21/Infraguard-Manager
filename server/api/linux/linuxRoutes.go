package linux

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
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
		return
	}
	out, err := sendCommandService(input)
	if err != nil {
		logger.Error("Error executing command on instance", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:sendCommand")
	c.JSON(http.StatusOK, out)
}

func executeScript(c *gin.Context) {
	logger.Info("IN:executeScript")

	// //Read form data
	// per := c.Request.FormValue("permission")
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	logger.Error("Error getting form data", err)
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }
	// content, _ := file.Open()
	// buf := bytes.NewBuffer(nil)
	// if _, err = io.Copy(buf, content); err != nil {
	// 	c.JSON(http.StatusExpectationFailed, err)
	// 	return
	// }
	var input model.Executable
	err := c.Bind(&input)
	if err != nil {
		logger.Error("error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	// input.Permission = per
	// input.Script = buf.Bytes()

	res, err := executeScriptService(input)
	if err != nil {
		logger.Error("Error executing script file", err)
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
