package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func getPublicAddressDB(serialId string) (model.Servers, error) {
	logger.Info("IN:getPublicAddress")

	res := model.Servers{}
	// query := "Select * from Servers where SerialID=?"
	gorm := db.MySqlConnection()
	if err := gorm.Table(db.ServerDB).Where("SerialID=?", serialId).Find(&res).Error; err != nil {
		logger.Error("Error getting activation details by SerialID", err)
		return res, err
	}
	logger.Info("IN:getPublicAddress")
	return res, nil
}
