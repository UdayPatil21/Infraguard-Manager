package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func getPublicAddressDB(machine_id string) (model.Servers, error) {
	logger.Info("IN:getPublicAddress")

	res := model.Servers{}
	// query := "Select * from Servers where InstanceID=?"
	gorm := db.MySqlConnection()
	if err := gorm.Table(db.ServerDB).Where("InstanceID=?", machine_id).Find(&res).Error; err != nil {
		logger.Error("Error getting activation details by InstanceID", err)
		return res, err
	}
	logger.Info("IN:getPublicAddress")
	return res, nil
}
