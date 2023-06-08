package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func getPublicAddressDB(machine_id string) (model.Servers, error) {
	logger.Info("IN:getPublicAddress")

	res := model.Servers{}
	// query := "Select * from agent where machine_id=?"
	gorm := db.MySqlConnection()
	// err := sql.QueryRow(query, machine_id).Scan(&res.Agent_id, &res.Name, &res.UserName, &res.MachineID, &res.PublicIP, &res.HostName, &res.OS, &res.CreatedAt)
	// if err != nil {
	// 	logger.Error("Error feching instance info", err)
	// 	return model.InstanceInfo{}, err
	// }
	// defer sql.Close()
	gorm.Table("Servers").Where("InstanceID=?", machine_id).Find(&res)
	logger.Info("IN:getPublicAddress")
	return res, nil
}
