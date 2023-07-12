package api

import (
	"bytes"
	"encoding/json"
	"errors"
	activation "infraguard-manager/api/agent-activation"
	"infraguard-manager/db"
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	"infraguard-manager/middleware/auth"
	model "infraguard-manager/models"
	"io/ioutil"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterInstance(c *gin.Context) {

	//Register all new instances
	var instanceInfo model.Agent
	response := model.Response{}
	//Statndard Output
	response.Status = false
	err := c.Bind(&instanceInfo)
	if err != nil {
		logger.Error("Error binding agent data", err)
		response.Error = err.Error()
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	//Check if agent is already resister
	if CheckAgentDB(instanceInfo) {
		str := "Agent Already Resistered"
		response.Data = str
		response.Status = true
		c.JSON(http.StatusOK, response)
		return
	}
	//validate activation details before register
	// if !validateAgentActivation(instanceInfo.AgentActivationID) {
	// 	logger.Error("Agent activation details not matched")
	// 	c.JSON(http.StatusExpectationFailed, "Agent activation details not matched")
	// 	return
	// }
	//Resister new server into the manager
	// err = ResisterInstanceService(instanceInfo)
	err = AgentService(instanceInfo)
	if err != nil {
		logger.Error("Error inserting instance info", err)
		response.Data = "Error in resistration of server"
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	response.Status = true
	response.Data = "Server Resister Successfully"
	c.JSON(http.StatusOK, response)

}

func AgentService(agent model.Agent) error {
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	// }
	agentBytes, err := json.Marshal(agent)
	if err != nil {
		logger.Error("Error marshling data", err)
		return err
	}
	jsonStr := string(agentBytes)
	//Get server URL from config
	base_url := configHelper.GetString("Infraguard-URL")
	//create req add neccessary headers
	client := &http.Client{}
	req, _ := http.NewRequest("POST", base_url+"/api/agent/servers", bytes.NewBufferString(jsonStr))
	// req.Header.Set("Authorization", configHelper.GetString("Authorization"))
	req.Header.Set("Authorization", auth.UserToken)
	req.Header.Set("Access-Infraguard", configHelper.GetString("Access-Infraguard"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending agent data to infraguard server", err)
		return err
	}
	//authoriaton
	//You do not have permission to access.
	//Check for permission error
	if resp.StatusCode == http.StatusForbidden {
		//Generate User Token and Try Again
		auth.GenerateUserToken(auth.LoginToken)
		//Resister Again with new User Token
		AgentService(agent)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		return err
	}
	logger.Info(string(respBytes))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("error in resistration of server")
	}

	// var out string
	// err = json.Unmarshal(respBytes, &out)
	// if err != nil {
	// 	logger.Error("Error unmarshling response data", err)
	// 	return err
	// }
	return nil
}

func CheckAgentDB(instance model.Agent) bool {
	logger.Info("IN:CheckAgentDB")
	server := model.Servers{}

	if result := db.DBInstance.Table(db.ServerDB).Where("InstanceID=?", instance.MachineID).Find(&server); result.Error != nil {
		logger.Error("Error getting activation details", result.Error)
		return false
	}
	if server.SerialID != "" {
		logger.Info("Agent already available")
		return true
	}
	logger.Info("OUT:CheckAgentDB")
	return false
}

type sh struct {
	service activation.ActivationService
}

func ServiceHandler() *sh {
	return &sh{}
}

//validate activation details before register
func validateAgentActivation(activationNumber int) bool {

	activation, err := activation.GetActivationByNumberDB(activationNumber)
	if err != nil {
		logger.Error("error getting activation data", err)
		return false
	}
	// id, _ := uuid.Parse(activationId)
	if activation.ID != activationNumber {
		return false
	}
	return true
}

//Update agent public ip
func UpdateAgent(c *gin.Context) {
	// gorm := db.MySqlConnection()
	server := model.UpdateServer{}
	new := model.Servers{}
	err := c.Bind(&server)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if err := db.DBInstance.Table(db.ServerDB).Where("InstanceID=?", server.InstanceID).Find(&new).Error; err != nil {
		logger.Error("Error getting data for updation", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	new.PublicIP = server.NetIP

	// d, _ := time.Parse("2006-01-02", new.CreatedDate)
	// new.CreatedDate = d.String()
	//Update servers set PubliIP=? where SerialID=?
	//Save updated data

	if result := db.DBInstance.Preloads(db.ServerDB).Table(db.ServerDB).Where("InstanceID=?", server.InstanceID).Update(&new); result.Error != nil {
		logger.Error("Error updating the activation details", result.Error)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

func UpdateServerInfo(c *gin.Context) {
	var instanceInfo model.Agent
	response := model.Response{}
	//Statndard Output
	response.Status = false
	err := c.Bind(&instanceInfo)
	if err != nil {
		logger.Error("Error binding agent data", err)
		response.Error = err.Error()
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	err = AgentService(instanceInfo)
	if err != nil {
		logger.Error("Error inserting instance info", err)
		response.Error = err.Error()
		c.JSON(http.StatusExpectationFailed, response)
		return
	}
	response.Data = "Server Info Updated Successfully"
	response.Status = true
	c.JSON(http.StatusOK, response)
}
