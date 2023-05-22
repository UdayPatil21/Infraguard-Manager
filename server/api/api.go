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
	}
	err = ResisterInstanceService(instanceInfo)
	if err != nil {
		logger.Error("Error inserting instance info", err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	c.JSON(http.StatusOK, "success")

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
	// id, err := insertResult.()
	// if err != nil {
	// 	logger.Error("impossible to retrieve last inserted id: %s", err)
	// 	return err
	// }
	// logger.Info("inserted id:", id)
	return nil
}
