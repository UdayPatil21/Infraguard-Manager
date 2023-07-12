package agent

import (
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExecuteScript(c *gin.Context) {
	logger.Info("IN:executeScript")

	var request model.Executable
	res := model.CmdOutput{}
	//Bind request data to the struct
	err := c.Bind(&request)
	if err != nil {
		logger.Error("error binding data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	//Standard Output
	res.Status = false
	//Check request data cannot be empty
	if request.Script == "" && request.SerialID == "" {
		res.Error = "Request Data Cannot Be Empty"
		c.JSON(http.StatusExpectationFailed, res)
		return
	}
	// res, err := executeScriptService(input)
	res, err = executeScriptService(request)
	if err != nil {
		logger.Error("Error executing script", err)
		// res.Error = "Error Executing Script Please Check"
		res.Error = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// s := SanitizeScript(res)
	logger.Info("OUT:executeScript")
	res.Status = true
	c.JSON(http.StatusOK, res)
}
func SanitizeScript(script string) string {
	s2 := strings.Replace(script, `\n`, "\n", -1)
	// s := strings.ReplaceAll(s2, "\\", "")
	return s2
}
