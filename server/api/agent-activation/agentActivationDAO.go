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
func addAgentActivationDB(data model.Activation) error {
	logger.Info("IN:getAgentActivationDB")

	//Generate activation_code (secret key)
	data.Actiovation_Code = hex.EncodeToString(([]byte(data.Activation_Name)))
	//Generate serial_id and activation_id
	data.Serial_Id = uuid.New()
	data.Activation_Id = uuid.New()

	//Assign created date
	data.Created_Date = time.Now()

	// Prepare insert query
	query := "insert into `activation` (`serial_id`,`activation_id`,`activation_code`,`activation_name`,`number_of_servers`,`created_by`,`created_date`) values(?,?,?,?,?,?,?)"
	sql := db.MySqlConnection()
	_, err := sql.Query(query, data.Serial_Id, data.Activation_Id, data.Actiovation_Code, data.Activation_Name, data.Server_Count, data.Created_By, data.Created_Date)
	if err != nil {
		logger.Error("Error inserting data to db", err)
		return err
	}
	logger.Info("OUT:getAgentActivationDB")
	return nil
}

//Get all activation data
func getAllActivationDB() ([]model.Activation, error) {
	logger.Info("IN:getAllActivationDB")
	activations := []model.Activation{}
	qry := "select * from activation"
	sql := db.MySqlConnection()
	row, err := sql.Query(qry)
	if err != nil {
		logger.Error("Error fetching activations")
		return activations, err
	}
	for row.Next() {
		data := model.Activation{}
		row.Scan(&data.Id, &data.Serial_Id, &data.Activation_Id, &data.Actiovation_Code, &data.Activation_Name, &data.Server_Count, &data.Created_By, &data.Created_Date)
		activations = append(activations, data)
	}

	logger.Info("OUT:getAllActivationDB")
	return activations, nil
}

type ActivationService interface {
	// getActivationByNameDB(string) (model.Activation, error)
	GetActivationByIdDB(uuid.UUID) (model.Activation, error)
}
type activation struct{}

// Get agent activation details by activation id
func (as *activation) GetActivationByIdDB(activationId uuid.UUID) (model.Activation, error) {
	logger.Info("IN:getActivationById")
	activation := model.Activation{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	query := "select * from `activation` where `activation_id`=?"
	sql := db.MySqlConnection()

	err := sql.QueryRow(query, activationId).Scan(&activation.Serial_Id, &activation.Activation_Id, &activation.Actiovation_Code, &activation.Activation_Name, &activation.Server_Count, &activation.Created_By, &activation.Created_Date)
	if err != nil {
		logger.Error("Error in getting agent data", err)
		return activation, err
	}
	logger.Info("OUT:getActivationById")
	return activation, nil
}

func GetActivationByIdDB(activationId uuid.UUID) (model.Activation, error) {
	logger.Info("IN:getActivationById")
	activation := model.Activation{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationId)
	query := "select * from `activation` where `activation_id`=?"
	sql := db.MySqlConnection()

	err := sql.QueryRow(query, activationId).Scan(&activation.Id, &activation.Serial_Id, &activation.Activation_Id, &activation.Actiovation_Code, &activation.Activation_Name, &activation.Server_Count, &activation.Created_By, &activation.Created_Date)
	if err != nil {
		logger.Error("Error in getting agent data", err)
		return activation, err
	}
	logger.Info("OUT:getActivationById")
	return activation, nil
}

// Get agent activation details by activation name
func GetActivationByNumberDB(activationNumber int) (model.Activation, error) {
	logger.Info("IN:getActivationById")
	activation := model.Activation{}
	//get uuid from activationId string
	//convert string to uuid
	// id, _ := uuid.Parse(activationName)
	query := "select * from `activation` where `id`=?"
	sql := db.MySqlConnection()

	err := sql.QueryRow(query, activationNumber).Scan(&activation.Id, &activation.Serial_Id, &activation.Activation_Id, &activation.Actiovation_Code, &activation.Activation_Name, &activation.Server_Count, &activation.Created_By, &activation.Created_Date)
	if err != nil {
		logger.Error("Error in getting agent data", err)
		return activation, err
	}
	logger.Info("OUT:getActivationById")
	return activation, nil
}

//Update specific activation data into the database
func updateAgentActivationDB(updateData model.Activation) error {
	logger.Info("IN:updateAgentActivationDB")
	sql := db.MySqlConnection()
	updateQuery := "update `activation` set `activation_name`=?,`number_of_servers`=? where `activation_id`=?"
	res, err := sql.Exec(updateQuery, updateData.Activation_Name, updateData.Server_Count)
	if err != nil {
		logger.Error("Error updating data", err)
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		logger.Error("Error executing query", err)
		return err
	}
	logger.Info("OUT:updateAgentActivationDB")
	return nil
}

//Delete agent activation data of specific id  from database
func DeleteAgentActivationByIdDB(activationId string) error {
	logger.Info("IN:DeleteAgentActivationByIdDB")
	//convert activationId into uuid
	id, _ := uuid.Parse(activationId)
	sql := db.MySqlConnection()
	//Prepare update query
	deleteQuery := "delete from `activation` where `activation_id`=?"
	res, err := sql.Exec(deleteQuery, id)
	if err != nil {
		logger.Error("Error in getting agent data", err)
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		logger.Error("Error in deleting the record", err)
		return err
	}
	logger.Info("OUT:DeleteAgentActivationByIdDB")
	return nil
}
