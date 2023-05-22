package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func getPublicAddressDB(agent_id string) (model.InstanceInfo, error) {
	logger.Info("IN:getPublicAddress")

	res := model.InstanceInfo{}
	query := "Select * from agent where agent_id=?"
	sql := db.MySqlConnection()
	result, err := sql.Query(query, agent_id)
	if err != nil {
		logger.Error("Error feching instance info", err)
		return model.InstanceInfo{}, err
	}
	result.Scan(&res)
	logger.Info("IN:getPublicAddress")
	return res, nil
}
