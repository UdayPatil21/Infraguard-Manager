package sheduler

import (
	"crypto/tls"
	"infraguard-manager/helpers/logger"
	"net/http"

	"github.com/robfig/cron"
)

// func CheckAgentStatus() {
// 	agents := getAllAgent()
// 	if len(agents) > 0 {
// 		for _, val := range agents {
// 			go Sheduler(val.PublicIP)
// 		}
// 	}
// }

// func getAllAgent() []model.InstanceInfo {
// 	agents := []model.InstanceInfo{}

// 	//Connect and fetch all active agents
// 	query := "select * from agent"
// 	sql := db.MySqlConnection()

// 	row, err := sql.Query(query)
// 	if err != nil {
// 		logger.Error("Error executing query", err)
// 		return agents
// 	}

// 	for row.Next() {
// 		var agent model.InstanceInfo
// 		err := row.Scan(&agent.Agent_id, &agent.Name, &agent.UserName, &agent.MachineID, &agent.PublicIP, &agent.HostName, &agent.OS, &agent.CreatedAt, &agent.Status)
// 		if err != nil {
// 			logger.Error("Error retriving data", err)
// 			return agents
// 		}
// 		agents = append(agents, agent)
// 	}
// 	return agents
// }

//Run cron job to check the status of the agent
func Sheduler(ip string) {
	c := cron.New()
	err := c.AddFunc("@every 5s", func() {
		logger.Info("Welcome")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
		client := &http.Client{Transport: tr}
		//send and execute command on the instance
		resp, err := client.Get(("http://" + /*strings.TrimSpace(ip)*/ "localhost" + ":8080/api/checkStatus"))
		if err != nil {
			logger.Error("Error")
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			logger.Info("Done")
		}
	})
	if err != nil {
		logger.Info("Error", err)
	}
	c.Start()
}
