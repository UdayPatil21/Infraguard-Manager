package activation

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Insert activation data into the database
func AddAgentActivation(c *gin.Context) {
	logger.Info("IN:AddAgentActivation")
	activationData := model.Clusters{}
	res := model.Response{}
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
	res.Data = "Activation Added Successfully"
	res.Status = true
	c.JSON(http.StatusOK, res)
}

//Get all activation data from the database
func GetAllActivation(c *gin.Context) {
	logger.Info("IN:GetAllActivation")
	res := model.Response{}
	activations, err := getAllActivationDB()
	if err != nil {
		logger.Error("Error getting info", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}

	if len(activations) > 0 {
		res.Data = activations
		res.Status = true
	} else {
		res.Data = "Activation Details Not Found"
		res.Status = false
	}
	logger.Info("OUT:GetAllActivation")
	c.JSON(http.StatusOK, res)
}

//Get activation details by id from database
func GetAgentActivationById(c *gin.Context) {
	logger.Info("IN:GetAgentActivationById")
	activationId := c.Param("id")
	res := model.Response{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	activation, err := GetActivationByIdDB(activationId)
	if err != nil {
		logger.Error("Error getting agent data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:GetAgentActivationById")
	if activation.ActivationID != "" {
		res.Data = activation
		res.Status = true
	} else {
		res.Data = "Activation Details Not Found"
		res.Status = false
	}

	c.JSON(http.StatusOK, res)
}

//Edit activation data and update into the database
func UpdateAgentActivation(c *gin.Context) {
	logger.Info("IN: UpdateAgentActivation")
	updateData := model.Clusters{}
	res := model.Response{}
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
	res.Data = "Activation Updated Successfully"
	res.Status = true
	c.JSON(http.StatusOK, res)
}

//Delete activation data from the database
func DeleteAgentActivationById(c *gin.Context) {
	logger.Info("IN:DeleteAgentActivation")
	res := model.Response{}
	id := c.Param("id")
	err := DeleteAgentActivationByIdDB(id)
	if err != nil {
		logger.Error("Error", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	logger.Info("OUT:DeleteAgentActivation")
	res.Data = "Activation Deleted Successfully"
	res.Status = true
	c.JSON(http.StatusOK, res)
}

//Get all activation data from the database
func GetAllServers(c *gin.Context) {
	logger.Info("IN:GetAllActivation")
	res := model.Response{}
	servers, err := getAllServersDB()
	if err != nil {
		logger.Error("Error getting info", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if len(servers) > 0 {
		res.Data = servers
		res.Status = true
	} else {
		res.Data = "Serves Details Not Found"
		res.Status = false
	}
	c.JSON(http.StatusOK, res)
}
