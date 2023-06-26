package api

import (
	"bytes"
	"encoding/json"
	"errors"
	activation "infraguard-manager/api/agent-activation"
	"infraguard-manager/db"
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
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
	// base_url := "https://9261-2401-4900-1f39-2814-bdc3-6488-72c0-9157.ngrok-free.app/api/agent/servers"
	// url := "https://e60a-2401-4900-1c5b-2c6a-2492-6f13-a798-2880.ngrok-free.app/user/get-data"
	// var v url.Values
	// v.Add("query", "a two level microprogram simulator")
	// v.Add("complete", "0")
	// v.Add("count", "10")
	// v.Add("model", "latest")
	// url := base_url + "?" + v.Encode()
	// client := &http.Client{Transport: tr}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", base_url, bytes.NewBufferString(jsonStr))
	req.Header.Set("Authorization", configHelper.GetString("Authorization"))
	req.Header.Set("Access-Infraguard", configHelper.GetString("Access-Infraguard"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// r, err := http.NewRequest("POST", base_url, bytes.NewBufferString(jsonStr))
	// if err != nil {
	// 	panic(err)
	// }
	// r.Header.Set("Authorization", "c87f8751-3899-4f50-9ef3-5239fd403120")
	// r.Header.Set("Access-Infraguard", "87d647c9-0dcd-4876-bdb9-3ec54e64dfa5")
	// r.Header.Add("Content-Type", "application/json")

	// resp, err := client.Post((configHelper.GetString("Infraguard-URL") /*"localhost"*/),
	// 	"application/json; charset=utf-8", bytes.NewBuffer(agentBytes))
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending agent data to infraguard server", err)
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		return err
	}
	logger.Info(string(respBytes))
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Error in resistration of server")
	}

	// var out string
	// err = json.Unmarshal(respBytes, &out)
	// if err != nil {
	// 	logger.Error("Error unmarshling response data", err)
	// 	return err
	// }
	return nil
}

//Insert instance data into the database
func ResisterInstanceService(info model.Servers) error {

	//query := "INSERT INTO `Servers` (`agent_id`,`name`, `user_name`, `machine_id`,`public_ip`,`hostname`,`os`,`created_at`,`status`,`activation_number`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	// `INSERT INTO `Servers` (`SerialID`,`Name`, `InstanceID`, `ServerID`,`ServerTags`,`PublicIP`,`PublicDNS`,`RegionID`,`OtherRegionName`,`Platform`,`OsVersion`,`ClusterID`,`InstanceProfileARN`,`Tags`,'AdditionalData',`ComputerName`,`ProviderID`,`ProjectID`,`CompanyID`,`MissingPatches`,`InstalledPatches`,`TotalPatches`,`AmiID`,`AmiCreationDetail`,
	// ``) VALUES (?,?,?,?,?,?,?,?,?,?)`
	// query := `INSERT INTO Servers (SerialID,Name,InstanceID,ServerTags,PublicIP,PublicDNS,RegionID,OtherRegionName,Platform,OsVersion,ClusterID,InstanceProfileARN,Tags,AdditionalData,ComputerName,ProviderID,ProjectID,CompanyID,
	// 	InstalledPatches,TotalPatches,AmiID,AmiCreationDetail,PatchCommandID,InstallingPatches,PatchInitiatedBy,PatchInstalledDate,IntervalsEmailDateTime,PatchScannedDate,SiteHostName,ResourceGroup,ResourceGroupID,SupportedAppsData,AgentActivationID,CreatedDate)
	// 	VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	gorm := db.MySqlConnection()
	// _, err := sql.Query(query, info.SerialID, info.Name, info.InstanceID, info.ServerTags, info.PublicIP, info.PublicDNS, info.RegionID, info.OtherRegionName, info.Platform, info.OsVersion, info.ClusterID, info.InstanceProfileARN, info.Tags, info.AdditionalData, info.ComputerName, info.ProviderID, info.ProjectID, info.CompanyID,
	// 	info.InstalledPatches, info.TotalPatches, info.AmiID, info.AmiCreationDetail, info.PatchCommandID, info.InstalledPatches, info.PatchInitiatedBy, info.PatchInstalledDate, info.IntervalsEmailDateTime, info.PatchScannedDate, info.SiteHostName, info.ResourceGroup, info.ResourceGroupID, info.SupportedAppsData, info.AgentActivationID, info.CreatedDate)
	// gorm.AutoMigrate(&info)

	//Insert into Servers () Values(&info)
	if err := gorm.Table(db.ServerDB).Create(&info).Error; err != nil {
		logger.Error("impossible insert AgentActivations: %s", err)
		return err
	}
	return nil
}

func CheckAgentDB(instance model.Agent) bool {
	logger.Info("IN:CheckAgentDB")
	// res := model.Servers{}
	server := model.Servers{}
	// query := "select * from Servers where  InstanceID=?"
	// sql := db.MySqlConnection()
	// err := sql.QueryRow(query, instance.InstanceID).Scan(
	// 	&res.ID,
	// 	&res.SerialID,
	// 	&res.Name,
	// 	&res.InstanceID,
	// 	&res.ServerID,
	// 	&res.InstanceType,
	// 	&res.ServerType,
	// 	&res.ServerTags,
	// 	&res.PublicIP,
	// 	&res.PublicDNS,
	// 	&res.PrivateIP,
	// 	&res.PrivateDNS,
	// 	&res.RegionID,
	// 	&res.OSID,
	// 	&res.OtherRegionName,
	// 	&res.AvailabilityZone,
	// 	&res.VPC,
	// 	&res.SubnetID,
	// 	&res.SecurityGroups,
	// 	&res.MappedDisks,
	// 	&res.ImageName,
	// 	&res.Platform,
	// 	&res.ADEnabled,
	// 	&res.OsVersion,
	// 	&res.ClusterID,
	// 	&res.InstanceProfileARN,
	// 	&res.Tags,
	// 	&res.AdditionalData,
	// 	&res.ComputerName,
	// 	&res.ProviderID,
	// 	&res.ProjectID,
	// 	&res.CompanyID,
	// 	&res.SSMStatus,
	// 	&res.IsServerRunning,
	// 	&res.IsServerLocked,
	// 	&res.MissingPatches,
	// 	&res.InstalledPatches,
	// 	&res.PatchDependenciesList,
	// 	&res.ComplianceStatus,
	// 	&res.TotalPatches,
	// 	&res.AmiID,
	// 	&res.AmiCreationDetail,
	// 	&res.IsPatchInstalled,
	// 	&res.PatchCommandID,
	// 	&res.InstallingPatches,
	// 	&res.PatchInitiatedBy,
	// 	&res.PatchInstalledDate,
	// 	&res.IntervalsEmailDateTime,
	// 	&res.PatchScannedDate,
	// 	&res.SiteHostName,
	// 	&res.ResourceGroup,
	// 	&res.ResourceGroupID,
	// 	&res.SupportedAppsData,
	// 	&res.IsMasterKeyAssigned,
	// 	&res.IsAbortedFromPolicy,
	// 	&res.LastHealthCheckAt,
	// 	&res.AgentActivationID,
	// 	&res.CreatedDate,
	// 	&res.IsActive,
	// 	&res.IsTerminated,
	// 	&res.IsDefault,
	// )
	// if err != nil {
	// 	logger.Error("Error retriving agent", err)
	// 	return false
	// }
	gorm := db.MySqlConnection()
	// gorm.AutoMigrate(&server)
	if result := gorm.Table(db.ServerDB).Where("InstanceID=?", instance.MachineID).Find(&server); result.Error != nil {
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
	gorm := db.MySqlConnection()
	server := model.UpdateServer{}
	new := model.Servers{}
	err := c.Bind(&server)
	if err != nil {
		logger.Error("Error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if err := gorm.Table(db.ServerDB).Where("InstanceID=?", server.InstanceID).Find(&new).Error; err != nil {
		logger.Error("Error getting data for updation", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	new.PublicIP = server.NetIP

	// d, _ := time.Parse("2006-01-02", new.CreatedDate)
	// new.CreatedDate = d.String()
	//Update servers set PubliIP=? where SerialID=?
	//Save updated data

	if result := gorm.Preloads(db.ServerDB).Table(db.ServerDB).Where("InstanceID=?", server.InstanceID).Update(&new); result.Error != nil {
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
