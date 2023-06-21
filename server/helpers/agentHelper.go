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
	checkStatus := model.CheckStatus{}
	err := c.Bind(&checkStatus)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, checkStatus)
		return
	}
	checkStatus.Status = false
	instanceInfo, err := linux.GetPublicAddressDB(checkStatus.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		c.JSON(http.StatusExpectationFailed, checkStatus)
		return
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	if err != nil {
		logger.Error("Error in unmarshaling", err)
		c.JSON(http.StatusExpectationFailed, checkStatus)
		return
	}
	resp, err := client.Get(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ + ":4200/api/checkStatus"))
	if err != nil {
		logger.Error("Error checking server status", err)
		c.JSON(http.StatusExpectationFailed, checkStatus)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusExpectationFailed, checkStatus)
		return
	}
	logger.Info("OUT:CheckStatus")
	checkStatus.Status = true
	c.JSON(http.StatusOK, checkStatus)
}
