package activation

import (
	"encoding/hex"
	"infraguard-manager/db"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"time"

	"github.com/google/uuid"
)

//
func addAgentActivationDB(data model.AgentActivations) error {
	logger.Info("IN:getAgentActivationDB")

	//Generate activation_code (secret key)
	data.ActivationCode = hex.EncodeToString(([]byte(data.ActivationName)))
	//Generate serial_id and activation_id
	data.SerialID = uuid.New().String()
	data.ActivationID = uuid.New().String()
	data.CreatedDateTime = time.Now()
	data.LastModifiedDateTime = time.Now()

	//Assign created date
	data.CreatedDateTime = time.Now()
	data.LastModifiedDateTime = time.Now()

	// Prepare insert query
	// query := "insert into `AgentActivations` (`SerialID`,`ActivationID`,`ActivationCode`,`ActivationName`,`TotalServers`,`CreatedBy`,`CreatedDateTime`,`ModifiedBy`,`LastModifiedDateTime`) values(?,?,?,?,?,?,?,?,?)"
	gorm := db.MySqlConnection()
	// _, err := sql.Query(query, data.SerialID, data.ActivationID, data.ActivationCode, data.ActivationName, data.TotalServers, data.CreatedBy, data.CreatedDateTime, data.ModifiedBy, data.LastModifiedDateTime)
	// if err != nil {
	// 	logger.Error("Error inserting data to db", err)
	// 	return err
	// }
	gorm.Table("AgentActivations").Create(&data)
	logger.Info("OUT:getAgentActivationDB")
	return nil
}

//Get all activation data
func getAllActivationDB() ([]model.AgentActivations, error) {
	logger.Info("IN:getAllActivationDB")
	activations := []model.AgentActivations{}
	// qry := "select * from AgentActivations"
	gorm := db.MySqlConnection()
	// row, err := sql.Query(qry)
	// if err != nil {
	// 	logger.Error("Error fetching activations")
	// 	return activations, err
	// }
	// for row.Next() {
	// 	data := model.AgentActivations{}
	// 	row.Scan(&data.ID, &data.SerialID, &data.ActivationID, &data.ActivationCode, &data.ActivationName, &data.TotalServers, &data.CreatedBy, &data.CreatedDateTime, &data.ModifiedBy, &data.LastModifiedDateTime, &data.IsActive)
	// 	activations = append(activations, data)
	// }
	gorm.Table("AgentActivations").Find(&activations)
	logger.Info("OUT:getAllActivationDB")
	return activations, nil
}

type ActivationService interface {
	// getActivationByNameDB(string) (model.Activation, error)
	GetActivationByIdDB(uuid.UUID) (model.AgentActivations, error)
}
type activation struct{}

// Get agent activation details by activation id
func (as *activation) GetActivationByIdDB(activationId uuid.UUID) (model.AgentActivations, error) {
	logger.Info("IN:getActivationById")
	activation := model.AgentActivations{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	// query := "select * from `AgentActivations` where `ActivationID`=?"
	gorm := db.MySqlConnection()

	// err := sql.QueryRow(query, activationId).Scan(&activation.ID, &activation.SerialID, &activation.ActivationID, &activation.ActivationCode, &activation.ActivationName, &activation.TotalServers, &activation.CreatedBy, &activation.CreatedDateTime, &activation.ModifiedBy, &activation.LastModifiedDateTime, &activation.IsActive)
	// if err != nil {
	// 	logger.Error("Error in getting agent data", err)
	// 	return activation, err
	// }
	gorm.Table("AgentActivations").Where("ActivationID=?", activationId).Find(&activation)
	logger.Info("OUT:getActivationById")
	return activation, nil
}

func GetActivationByIdDB(activationId string) (model.AgentActivations, error) {
	logger.Info("IN:getActivationById")
	activation := model.AgentActivations{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	// query := "select * from `AgentActivations` where `ActivationID`=?"
	gorm := db.MySqlConnection()

	// err := sql.QueryRow(query, activationId).Scan(&activation.ID, &activation.SerialID, &activation.ActivationID, &activation.ActivationCode, &activation.ActivationName, &activation.TotalServers, &activation.CreatedBy, &activation.CreatedDateTime, &activation.ModifiedBy, &activation.LastModifiedDateTime, &activation.IsActive)
	// if err != nil {
	// 	logger.Error("Error in getting agent data", err)
	// 	return activation, err
	// }
	gorm.Table("AgentActivations").Where("ActivationID=?", activationId).Find(&activation)
	logger.Info("OUT:getActivationById")
	return activation, nil
}

// Get agent activation details by activation name
func GetActivationByNumberDB(activationNumber int) (model.AgentActivations, error) {
	logger.Info("IN:getActivationById")
	activation := model.AgentActivations{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationName)
	gorm := db.MySqlConnection()
	// query := "select * from AgentActivations where ID=?"
	// err := sql.QueryRow(query, activationNumber).Scan(&activation.ID, &activation.SerialID, &activation.ActivationID, &activation.ActivationCode, &activation.ActivationName, &activation.TotalServers, &activation.CreatedBy, &activation.CreatedDateTime, &activation.ModifiedBy, &activation.LastModifiedDateTime, &activation.IsActive)
	gorm.Table("AgentActivations").Where("ID=?", activationNumber).Find(&activation)
	// if err != nil {
	// 	logger.Error("Error in getting agent data")
	// 	return activation, err
	// }

	logger.Info("OUT:getActivationById")
	return activation, nil
}

//Update specific activation data into the database
func updateAgentActivationDB(updateData model.AgentActivations) error {
	logger.Info("IN:updateAgentActivationDB")
	gorm := db.MySqlConnection()
	new := model.AgentActivations{}
	// updateQuery := "update `AgentActivations` set `ActivationName`=?,`TotalServers`=? where `ActivationID`=?"
	// res, err := sql.Exec(updateQuery, updateData.ActivationName, updateData.TotalServers)
	// if err != nil {
	// 	logger.Error("Error updating data", err)
	// 	return err
	// }
	// count, err := res.RowsAffected()
	// if count == 0 {
	// 	logger.Error("Error executing query", err)
	// 	return err
	// }
	//First fetch data you want to update
	gorm.Table("AgentActivations").Where("ActivationID=?", updateData.ActivationID).Find(&new)
	//change the required field
	if updateData.SerialID != "" {
		new.SerialID = updateData.SerialID
	}
	if updateData.ActivationCode != "" {
		new.ActivationCode = updateData.ActivationCode
	}
	if updateData.ActivationName != "" {
		new.ActivationName = updateData.ActivationName
	}
	if updateData.TotalServers != 0 {
		new.TotalServers = updateData.TotalServers
	}
	if updateData.ModifiedBy != 0 {
		new.ModifiedBy = updateData.ModifiedBy
	}
	if updateData.IsActive != "" {
		new.IsActive = updateData.IsActive
	}
	new.LastModifiedDateTime = time.Now()

	//Save updated data
	gorm.Table("AgentActivations").Where("ActivationID=?", updateData.ActivationID).Update(&new)
	logger.Info("OUT:updateAgentActivationDB")
	return nil
}

//Delete agent activation data of specific id  from database
func DeleteAgentActivationByIdDB(activationId string) error {
	logger.Info("IN:DeleteAgentActivationByIdDB")
	//convert activationId into uuid
	// id, _ := uuid.Parse(activationId)
	activation := model.AgentActivations{}
	gorm := db.MySqlConnection()
	//Prepare update query
	// deleteQuery := "delete from `activation` where `activation_id`=?"
	// res, err := sql.Exec(deleteQuery, id)
	// if err != nil {
	// 	logger.Error("Error in getting agent data", err)
	// 	return err
	// }
	// count, err := res.RowsAffected()
	// if count == 0 {
	// 	logger.Error("Error in deleting the record", err)
	// 	return err
	// }
	gorm.Table("AgentActivations").Where("ActivationID = ?", activationId).Delete(&activation)
	logger.Info("OUT:DeleteAgentActivationByIdDB")
	return nil
}
