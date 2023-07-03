package helper

import (
	"crypto/tls"
	"infraguard-manager/api/linux"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Check Agent Status
func CheckStatus(c *gin.Context) {
	logger.Info("IN:CheckStatus")
	request := model.CheckStatus{}
	response := model.Response{}

	err := c.Bind(&request)
	if err != nil {
		logger.Error("Error binding data", err)
		response.Error = err
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	request.Status = false
	//Standard Output
	response.Status = false
	response.Data = request.SerialID
	//Check if serialID present or not
	if request.SerialID == "" {
		response.Error = "Please Enter Server ID"
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	instanceInfo, err := linux.GetPublicAddressDB(request.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		response.Error = "Provide Correct Server ID"
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	if err != nil {
		logger.Error("Error in unmarshaling", err)
		response.Error = err
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	resp, err := client.Get(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ + ":4200/api/checkStatus"))
	if err != nil {
		logger.Error("Error checking server status", err)
		response.Error = err
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	logger.Info("OUT:CheckStatus")
	response.Status = true
	c.JSON(http.StatusOK, response)
}
