package helper

import (
	"crypto/tls"
	"infraguard-manager/helpers/logger"
	"io/ioutil"
	"net/http"
)

func CheckStatus(IP string) bool {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	//check status of agent server
	resp, err := client.Get(("http://" + /*strings.TrimSpace(IP)*/ "localhost" + ":8080/api/checkStatus"))
	if err != nil {
		logger.Error("Error:", err)
		return false
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error", err)
	}
	if resp.StatusCode != 200 && string(responseData) != "success" {
		return false
	}
	return true
}
