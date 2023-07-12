package agent

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

type Request struct {
	Script   string `json:"Script"`
	Platform string `json:"Platform"`
}

// Execute scripts
func executeScriptService(input model.Executable) (model.CmdOutput, error) {
	logger.Info("IN:executeScriptService")
	request := Request{}
	cmd := model.CmdOutput{}
	instanceInfo, err := GetPublicAddressDB(input.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return cmd, err
	}
	//Trim public ip
	//Add server platform to request
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	request.Script = input.Script
	request.Platform = instanceInfo.Platform
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	// }
	// client := &http.Client{Transport: tr}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		logger.Error("Error unmarshaling script", err)
		return cmd, err
	}
	jsonStr := string(reqBytes)

	// resp, err := client.Post(("http://" + strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ + ":4200/api/linux/script/execute"),
	// 	"application/json; charset=utf-8", bytes.NewBuffer(scriptByte))
	// if err != nil {
	// 	logger.Error("Error executing script file on instance", err)
	// 	return cmd, err
	// }

	//create http client request and execute scripts
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://"+strings.TrimSpace(instanceInfo.PublicIP) /*"localhost"*/ +":4200/api/script/execute", bytes.NewBufferString(jsonStr))
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
	request := Request{}
	instanceInfo, err := GetPublicAddressDB(input.SerialID)
	if err != nil {
		logger.Error("Error getting instance info from DB", err)
		return cmd, err
	}
	//Trim public ip
	//Add server platform to request
	instanceInfo.PublicIP = strings.TrimSpace(instanceInfo.PublicIP)
	request.Script = input.Script
	request.Platform = instanceInfo.Platform

	reqBytes, err := json.Marshal(request)
	if err != nil {
		logger.Error("Error unmarshaling script", err)
		return cmd, err
	}
	jsonStr := string(reqBytes)
	//create http client request and execute scripts
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://"+"localhost"+":4200/api/script/execute", bytes.NewBufferString(jsonStr))
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
