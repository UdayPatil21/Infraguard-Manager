package api

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterInstance(c *gin.Context) {
	//Register all new instances
	var instanceInfo model.InstanceInfo

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
func ResisterInstanceService(info model.InstanceInfo) error {
	agent_id := uuid.New()
	query := "INSERT INTO `agent` (`agent_id`,`name`, `user_name`, `machine_id`,`public_ip`,`hostname`,`os`,`created_at`) VALUES (?,?,?,?,?,?,?,?)"
	sql := db.MySqlConnection()
	_, err := sql.Query(query, agent_id, info.Name, info.UserName, info.MachineID, info.PublicIP, info.HostName, info.OS, info.CreatedAt)
	if err != nil {
		logger.Error("impossible insert agent: %s", err)
		return err
	}
	defer sql.Close()
	// id, err := insertResult.()
	// if err != nil {
	// 	logger.Error("impossible to retrieve last inserted id: %s", err)
	// 	return err
	// }
	// logger.Info("inserted id:", id)
	return nil
}

func CheckAgentDB(instance model.InstanceInfo) bool {
	logger.Info("IN:CheckAgentDB")
	res := model.InstanceInfo{}
	query := "select * from agent where  machine_id=? and public_ip=?"
	sql := db.MySqlConnection()
	err := sql.QueryRow(query, instance.MachineID, instance.PublicIP).Scan(&res.Agent_id, &res.Name, &res.UserName, &res.MachineID, &res.PublicIP, &res.HostName, &res.OS, &res.CreatedAt)
	if err != nil {
		logger.Error("Error retriving agent", err)
		return false
	}
	defer sql.Close()
	if res.Agent_id != uuid.Nil {
		logger.Info("Agent already available")
		return true
	}
	logger.Info("OUT:CheckAgentDB")
	return false
}
