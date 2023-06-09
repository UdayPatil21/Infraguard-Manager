package activation

import (
	"encoding/hex"
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"time"

	"github.com/google/uuid"
)

func addAgentActivationDB(data model.Clusters) error {
	logger.Log.Info("IN:getAgentActivationDB")

	//Generate activation_code (secret key)
	data.ActivationCode = hex.EncodeToString(([]byte(data.Name)))
	//Generate serial_id and activation_id
	data.SerialID = uuid.New().String()
	data.ActivationID = uuid.New().String()

	//Assign created date
	data.CreatedDate = time.Now()
	data.ModifiedDate = time.Now()

	// Prepare insert query
	// query := "insert into `AgentActivations` (`SerialID`,`ActivationID`,`ActivationCode`,`ActivationName`,`TotalServers`,`CreatedBy`,`CreatedDateTime`,`ModifiedBy`,`LastModifiedDateTime`) values(?,?,?,?,?,?,?,?,?)"
	// gorm := db.MySqlConnection()
	// _, err := sql.Query(query, data.SerialID, data.ActivationID, data.ActivationCode, data.ActivationName, data.TotalServers, data.CreatedBy, data.CreatedDateTime, data.ModifiedBy, data.LastModifiedDateTime)
	// if err != nil {
	// 	logger.Log.Sugar().Errorf("Error inserting data to db", err)
	// 	return err
	// }
	if err := db.DBInstance.Table(db.ActivationDB).Create(&data).Error; err != nil {
		logger.Log.Sugar().Errorf("Error inserting data", err)
		return err
	}
	logger.Log.Info("OUT:getAgentActivationDB")
	return nil
}

// Get all activation data
func getAllActivationDB() ([]model.Clusters, error) {
	logger.Log.Info("IN:getAllActivationDB")
	activations := []model.Clusters{}
	// qry := "select * from AgentActivations"
	// gorm := db.MySqlConnection()
	if err := db.DBInstance.Table(db.ActivationDB).Find(&activations).Error; err != nil {
		logger.Log.Sugar().Errorf("Error getting all the activation details", err)
		return activations, err
	}
	logger.Log.Info("OUT:getAllActivationDB")
	return activations, nil
}

type ActivationService interface {
	// getActivationByNameDB(string) (model.Activation, error)
	GetActivationByIdDB(uuid.UUID) (model.Clusters, error)
}
type activation struct{}

// Get agent activation details by activation id
func (as *activation) GetActivationByIdDB(activationId uuid.UUID) (model.Clusters, error) {
	logger.Log.Info("IN:getActivationById")
	activation := model.Clusters{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	// query := "select * from `AgentActivations` where `ActivationID`=?"
	// gorm := db.MySqlConnection()

	// err := sql.QueryRow(query, activationId).Scan(&activation.ID, &activation.SerialID, &activation.ActivationID, &activation.ActivationCode, &activation.ActivationName, &activation.TotalServers, &activation.CreatedBy, &activation.CreatedDateTime, &activation.ModifiedBy, &activation.LastModifiedDateTime, &activation.IsActive)
	// if err != nil {
	// 	logger.Log.Sugar().Errorf("Error in getting agent data", err)
	// 	return activation, err
	// }
	db.DBInstance.Table("AgentActivations").Where("ActivationID=?", activationId).Find(&activation)
	logger.Log.Info("OUT:getActivationById")
	return activation, nil
}

func GetActivationByIdDB(activationId string) (model.Clusters, error) {
	logger.Log.Info("IN:GetActivationByIdDB")
	activation := model.Clusters{}
	// query := "select * from `AgentActivations` where `ActivationID`=?"
	// gorm := db.MySqlConnection()
	if err := db.DBInstance.Table(db.ActivationDB).Where("ActivationID=?", activationId).Find(&activation).Error; err != nil {
		logger.Log.Sugar().Errorf("Error getting activation", err)
		return activation, err
	}
	logger.Log.Info("OUT:GetActivationByIdDB")
	return activation, nil
}

// Get agent activation details by activation name
func GetActivationByNumberDB(activationNumber int) (model.Clusters, error) {
	logger.Log.Info("IN:GetActivationByNumberDB")
	activation := model.Clusters{}
	// gorm := db.MySqlConnection()
	//Select * from AgentActivations where ID=?
	if err := db.DBInstance.Table(db.ActivationDB).Where("ID=?", activationNumber).Find(&activation).Error; err != nil {
		logger.Log.Sugar().Errorf("Error in getting agent data")
		return activation, err
	}

	logger.Log.Info("OUT:GetActivationByNumberDB")
	return activation, nil
}

// Update specific activation data into the database
func updateAgentActivationDB(updateData model.Clusters) error {
	logger.Log.Info("IN:updateAgentActivationDB")
	// gorm := db.MySqlConnection()
	new := model.Clusters{}
	// updateQuery := "update `AgentActivations` set `ActivationName`=?,`TotalServers`=? where `ActivationID`=?"
	//First fetch data you want to update
	if err := db.DBInstance.Table(db.ActivationDB).Where("ActivationID=?", updateData.ActivationID).Find(&new).Error; err != nil {
		logger.Log.Sugar().Errorf("Error getting data for updation", err)
		return err
	}
	//change the required field
	// if updateData.SerialID != "" {
	// 	new.SerialID = updateData.SerialID
	// }
	// if updateData.ActivationCode != "" {
	// 	new.ActivationCode = updateData.ActivationCode
	// }
	// if updateData.ActivationName != "" {
	// 	new.ActivationName = updateData.ActivationName
	// }
	// if updateData.TotalServers != 0 {
	// 	new.TotalServers = updateData.TotalServers
	// }
	// if updateData.ModifiedBy != 0 {
	// 	new.ModifiedBy = updateData.ModifiedBy
	// }
	// if updateData.IsActive != "" {
	// 	new.IsActive = updateData.IsActive
	// }
	// new.LastModifiedDateTime = time.Now()

	//Save updated data
	if result := db.DBInstance.Table(db.ActivationDB).Where("ActivationID=?", updateData.ActivationID).Update(&new); result.Error != nil {
		logger.Log.Sugar().Errorf("Error updating the activation details", result.Error)
		return result.Error
	}
	logger.Log.Info("OUT:updateAgentActivationDB")
	return nil
}

// Delete agent activation data of specific id  from database
func DeleteAgentActivationByIdDB(activationId string) error {
	logger.Log.Info("IN:DeleteAgentActivationByIdDB")
	activation := model.Clusters{}
	// gorm := db.MySqlConnection()
	// deleteQuery := "delete from `activation` where `activation_id`=?"
	if err := db.DBInstance.Table(db.ActivationDB).Where("ActivationID = ?", activationId).Delete(&activation).Error; err != nil {
		logger.Log.Sugar().Errorf("Error in deletion of activation", err)
		return err
	}
	logger.Log.Info("OUT:DeleteAgentActivationByIdDB")
	return nil
}

func GetAllServersDB() ([]model.Servers, error) {
	logger.Log.Info("IN:GetAllServersDB")
	servers := []model.Servers{}
	// cluster.SerialID AS ActivationSerialID,
	// cluster.Name AS ActivationName,
	// cluster.ActivationID,
	// cluster.ActivationCode,
	// cluster.TotalServers,
	// servers.SerialID AS ServerSerialID,
	// servers.Name AS ServerName,
	// servers.InstanceID,
	// servers.ServerID,
	// servers.PublicIP,
	// servers.PrivateIP,
	// servers.Platform,
	// servers.ComputerName,
	// servers.SiteHostName
	qry := `
	SELECT
		*
	FROM
		infraguard_manager.Clusters cluster
			INNER JOIN
		infraguard_manager.Servers servers ON cluster.ID = servers.AgentActivationID
	WHERE
		cluster.IsActive = 'Yes'
			AND cluster.IsDeleted = 'No'
			AND servers.IsActive = 'Yes'
			AND servers.IsTerminated = 'No'
			AND cluster.ProviderID = 5;`

	// gorm := db.MySqlConnection()

	// gorm.Table(db.ServerDB).Table(db.ActivationDB).Joins("JOIN Clusters ON Clusters.ID = Servers.AgentActivationID").Where().Find(&servers)
	if err := db.DBInstance.Raw(qry).Scan(&servers).Error; err != nil {
		logger.Log.Sugar().Errorf("Error getting all the activation details", err)
		return servers, err
	}
	logger.Log.Info("OUT:GetAllServersDB")
	return servers, nil
}
