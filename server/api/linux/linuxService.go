package linux

import (
	"bytes"
	"encoding/json"
	"errors"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io/ioutil"
	"net/http"
	"strings"
)

//Execute commands
func sendCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sendCommandService")
	//Get public ip from db
	instanceInfo, err := GetPublicAddressDB(input.MachineID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return nil, err
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	jsonReq, _ := json.Marshal(input)

	//create http client request and execute commands
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://"+strings.TrimSpace(instanceInfo.PublicIP)+":4200/api/linux/command/execute", bytes.NewBuffer(jsonReq))
	req.Header.Set("Authorization", model.TokenString)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error executing command", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error executing command")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error:", err)
		return "", err
	}
	logger.Info("OUT:sendCommandService")
	return string(responseData), nil
}

// Execute scripts
func executeScriptService(input model.Executable) (model.CmdOutput, error) {
	logger.Info("IN:executeScriptService")
	// marshal request data
	// jsonReq, _ := json.Marshal(input)
	cmd := model.CmdOutput{}
	instanceInfo, err := GetPublicAddressDB(input.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return cmd, err
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	// }
	// client := &http.Client{Transport: tr}
	scriptByte, err := json.Marshal(input.Script)
	if err != nil {
		logger.Error("Error unmarshaling script", err)
		return cmd, err
	}
	// resp, err := client.Post(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ + ":4200/api/linux/script/execute"),
	// 	"application/json; charset=utf-8", bytes.NewBuffer(scriptByte))
	// if err != nil {
	// 	logger.Error("Error executing script file on instance", err)
	// 	return cmd, err
	// }
	//create http client request and execute scripts
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://"+strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ +":4200/api/linux/script/execute", bytes.NewBuffer(scriptByte))
	req.Header.Set("Authorization", model.TokenString)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error executing script file on instance", err)
		return cmd, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return cmd, errors.New("error executing script")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		return cmd, err
	}

	//Convert response data into the object
	err = json.Unmarshal(responseData, &cmd.Output)
	if err != nil {
		logger.Error("Error converting output", err)
		return cmd, err
	}
	logger.Info("OUT:executeScriptService")
	return cmd, nil
}

// Execute scripts
func executeScriptLocal(input model.Executable) (model.CmdOutput, error) {
	logger.Info("IN:executeScriptService")
	cmd := model.CmdOutput{}
	instanceInfo, err := GetPublicAddressDB(input.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return cmd, err
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)

	scriptByte, err := json.Marshal(input.Script)
	if err != nil {
		logger.Error("Error unmarshaling script", err)
		return cmd, err
	}
	//create http client request and execute scripts
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://"+"localhost"+":4200/api/linux/script/execute", bytes.NewBuffer(scriptByte))
	req.Header.Set("Authorization", model.TokenString)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error executing script file on instance", err)
		cmd.Output = resp.Status
		return cmd, err
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		cmd.Output = resp.Status
		return cmd, err
	}
	if resp.StatusCode != http.StatusOK {
		cmd.Output = resp.Status
		return cmd, errors.New("error executing script")
	}
	//Convert response data into the object
	err = json.Unmarshal(responseData, &cmd.Output)
	if err != nil {
		logger.Error("Error converting output", err)
		return cmd, err
	}
	logger.Info("OUT:executeScriptService")
	return cmd, nil
}
