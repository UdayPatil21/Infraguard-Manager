package activation

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
)

func addAgentActivationService(data model.Clusters) error {
	logger.Log.Info("IN: addAgentActivationService")
	//Insert activation data into the table
	err := addAgentActivationDB(data)
	if err != nil {
		logger.Log.Sugar().Errorf("Error in inseting data", err)
		return err
	}
	logger.Log.Info("OUT:addAgentActivationService")
	return nil
}
