package linux

import (
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func GetPublicAddressDB(serialId string) (model.Servers, error) {
	logger.Info("IN:getPublicAddress")

	res := model.Servers{}
	// query := "Select * from Servers where SerialID=?"
	gorm := db.MySqlConnection()
	if err := gorm.Table(db.ServerDB).Where("SerialID=? AND IsActive=? AND IsTerminated=?", serialId, "Yes", "No").Find(&res).Error; err != nil {
		logger.Error("Error getting server details by SerialID", err)
		return res, err
	}
	// defer gorm.Close()
	logger.Info("IN:getPublicAddress")
	return res, nil
}
