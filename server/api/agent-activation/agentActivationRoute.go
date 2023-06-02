package activation

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Insert activation data into the database
func AddAgentActivation(c *gin.Context) {
	logger.Info("IN:AddAgentActivation")
	activationData := model.Activation{}
	err := c.Bind(&activationData)
	if err != nil {
		logger.Error("Error in binding the activation data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	err = addAgentActivationService(activationData)
	if err != nil {
		logger.Error("Error in addAgentActivationService", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:AddAgentActivation")
	c.JSON(http.StatusOK, "success")
}

//Get all activation data from the database
func GetAllActivation(c *gin.Context) {
	logger.Info("IN:GetAllActivation")
	activations, err := getAllActivationDB()
	if err != nil {
		logger.Error("Error getting info", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:GetAllActivation")
	c.JSON(http.StatusOK, activations)
}

//Get activation details by id from database
func GetAgentActivationById(c *gin.Context) {
	logger.Info("IN:GetAgentActivationById")
	activationId := c.Param("id")
	//get uuid from activationId string
	//convert string to uuid
	id, _ := uuid.Parse(activationId)
	activation, err := GetActivationByIdDB(id)
	if err != nil {
		logger.Error("Error getting agent data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:GetAgentActivationById")
	c.JSON(http.StatusOK, activation)
}

//Edit activation data and update into the database
func UpdateAgentActivation(c *gin.Context) {
	logger.Info("IN: UpdateAgentActivation")
	updateData := model.Activation{}
	err := c.Bind(&updateData)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	err = updateAgentActivationDB(updateData)
	if err != nil {
		logger.Error("Error updating record", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT: UpdateAgentActivation")
	c.JSON(http.StatusOK, "success")
}

//Delete activation data from the database
func DeleteAgentActivationById(c *gin.Context) {
	logger.Info("IN:DeleteAgentActivation")
	id := c.Param("id")
	err := DeleteAgentActivationByIdDB(id)
	if err != nil {
		logger.Error("Error", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:DeleteAgentActivation")
	c.JSON(http.StatusOK, "success")
}
