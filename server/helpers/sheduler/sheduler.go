package sheduler

import (
	"crypto/tls"
	"infraguard-manager/db"
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/robfig/cron"
)

//Cron to
//Fetch servers data from database
//Update server status on infraguard database
func Scheduler() {
	c := cron.New()
	servers := []model.Servers{}
	err := c.AddFunc("@every 15s", func() {
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
			AND cluster.ProviderID = 5;
		`
		//Fetch all active servers data from the database
		gorm := db.MySqlConnection()
		// gorm.Table(db.ServerDB).Table(db.ActivationDB).Joins("JOIN Clusters ON Clusters.ID = Servers.AgentActivationID").Where().Find(&servers)
		if err := gorm.Raw(qry).Scan(&servers).Error; err != nil {
			logger.Error("Error getting all the servers details", err)
		}
		//check server status and update into the database
		//create seperate goroutines for every server
		if len(servers) > 0 {
			for _, val := range servers {
				go CheckStatus(val.SerialID, val.PublicIP)
			}
		}
		//We can stop cron in there is no server to check the status
		// else {
		// 	c.Stop()
		// }
	})
	if err != nil {
		logger.Info("Error", err)
	}
	c.Start()
}

//Run cron job to check the status of the agent
func CheckStatus(SerialID, PublicIP string) {
	PublicIP = strings.TrimSpace(PublicIP)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(("http://" + /*strings.TrimSpace(PublicIP)*/ "localhost" + ":4200/api/checkStatus"))
	if err != nil {
		logger.Error("Error checking server status", err)
		ChangeServerStatus(SerialID)
	}
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			ChangeServerStatus(SerialID)
		}
	}

}
func ChangeServerStatus(SerialID string) {
	logger.Info("IN:ChangeServerStatus")
	//Get server URL from config
	base_url := configHelper.GetString("Infraguard-URL")
	//create req add neccessary headers
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", base_url+"/api/agent/servers/"+SerialID, nil)
	req.Header.Set("Authorization", configHelper.GetString("Authorization"))
	req.Header.Set("Access-Infraguard", configHelper.GetString("Access-Infraguard"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error deleting agent server", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
	}
	logger.Info(string(respBytes))
	defer resp.Body.Close()
	logger.Info("OUT:ChangeServerStatus")
}
