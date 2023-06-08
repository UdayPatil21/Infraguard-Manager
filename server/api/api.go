package api

import (
	activation "infraguard-manager/api/agent-activation"
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterInstance(c *gin.Context) {
	//Register all new instances
	var instanceInfo model.Servers

	err := c.Bind(&instanceInfo)
	if err != nil {
		logger.Error("Error binding agent data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	//Check if agent is already resister
	if CheckAgentDB(instanceInfo) {
		c.JSON(http.StatusOK, "Agent Already Resistered")
		return
	}
	//validate activation details before register
	if !validateAgentActivation(instanceInfo.AgentActivationID) {
		logger.Error("Agent activation details not matched")
		c.JSON(http.StatusExpectationFailed, "Agent activation details not matched")
		return
	}
	//Resister new server into the manager
	err = ResisterInstanceService(instanceInfo)
	if err != nil {
		logger.Error("Error inserting instance info", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, "Success")

}

//Insert instance data into the database
func ResisterInstanceService(info model.Servers) error {
	// agent_id := uuid.New()
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
	gorm.Table("Servers").Create(&info)
	// if err != nil {
	// 	logger.Error("impossible insert AgentActivations: %s", err)
	// 	return err
	// }
	defer gorm.Close()
	// id, err := insertResult.()
	// if err != nil {
	// 	logger.Error("impossible to retrieve last inserted id: %s", err)
	// 	return err
	// }
	// logger.Info("inserted id:", id)
	return nil
}

func CheckAgentDB(instance model.Servers) bool {
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
	gorm.AutoMigrate(&server)
	gorm.Table("Servers").Where("InstanceID=?", instance.InstanceID).Find(&server)
	// json.Unmarshal(serverData, &res)
	defer gorm.Close()
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
