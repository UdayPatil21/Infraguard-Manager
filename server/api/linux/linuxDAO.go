package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func getPublicAddressDB(machine_id string) (model.InstanceInfo, error) {
	logger.Info("IN:getPublicAddress")

	res := model.InstanceInfo{}
	query := "Select * from agent where machine_id=?"
	sql := db.MySqlConnection()
	err := sql.QueryRow(query, machine_id).Scan(&res.Agent_id, &res.Name, &res.UserName, &res.MachineID, &res.PublicIP, &res.HostName, &res.OS, &res.CreatedAt)
	if err != nil {
		logger.Error("Error feching instance info", err)
		return model.InstanceInfo{}, err
	}
	defer sql.Close()
	// result.Scan(&res)
	logger.Info("IN:getPublicAddress")
	return res, nil
}
