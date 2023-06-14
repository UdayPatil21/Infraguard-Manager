package activation

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func addAgentActivationService(data model.Clusters) error {
	logger.Info("IN: addAgentActivationService")
	//Insert activation data into the table
	err := addAgentActivationDB(data)
	if err != nil {
		logger.Error("Error in inseting data", err)
		return err
	}
	logger.Info("OUT:addAgentActivationService")
	return nil
}
