package linux

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Execute commands
func sendCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sendCommandService")
	//Get public ip from db
	instanceInfo, err := getPublicAddressDB(input.MachineID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return nil, err
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	jsonReq, _ := json.Marshal(input)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	//send and execute command on the instance
	resp, err := client.Post(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost" */ + ":4200/api/linux/send-command"),
		"application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Error("Error executing command", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Error Executing Command")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error:", err)
		return "", err
	}
	logger.Info("OUT:sendCommandService")
	return string(responseData), nil
}

//Execute sudo commands
func sudoCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sudoCommandService")
	//Get public ip from db
	instanceInfo, err := getPublicAddressDB(input.MachineID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return nil, err
	}
	_ = instanceInfo

	jsonReq, _ := json.Marshal(input)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	//send and execute command on the instance
	resp, err := client.Post(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*localhost" */ + ":4200/api/linux/sudo-command"),
		"application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Error("Error executing command", err)
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("OUT:sudoCommandService")
	return string(responseData), nil
}

//Execute scripts
func executeScriptService(input model.Executable) (any, error) {
	logger.Info("IN:executeScriptService")
	// marshal request data
	// jsonReq, _ := json.Marshal(input)

	instanceInfo, err := getPublicAddressDB(input.MachineID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return nil, err
	}
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ + ":4200/api/linux/execute-script"),
		"application/json; charset=utf-8", bytes.NewBuffer(input.Script))
	if err != nil {
		logger.Error("Error executing script file on instance", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Error Executing Script")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		return nil, err
	}
	logger.Info("OUT:executeScriptService")
	return string(responseData), nil
}
