package linux

// func InitLinuxRoutes(routeGroup *gin.RouterGroup) {
// 	r := routeGroup.Group("/platform/linux").Use(middleware.ValidateDomain())
// 	r.POST("/script/command", executeCommand)
// 	// r.POST("/script/execute", executeScript)
// }

// func executeCommand(c *gin.Context) {
// 	logger.Log.Info("IN:sendCommand")
// 	input := model.RunCommand{}
// 	err := c.Bind(&input)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error binding data", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 		return
// 	}
// 	out, err := sendCommandService(input)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error executing command on instance", err)
// 		c.JSON(http.StatusExpectationFailed, err.Error())
// 		return
// 	}
// 	logger.Log.Info("OUT:sendCommand")
// 	c.JSON(http.StatusOK, out)
// }

// func executeScript(c *gin.Context) {
// 	logger.Log.Info("IN:executeScript")

// 	var request model.Executable
// 	res := model.CmdOutput{}
// 	//Bind request data to the struct
// 	err := c.Bind(&request)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("error binding data", err)
// 		c.JSON(http.StatusExpectationFailed, err)
// 		return
// 	}
// 	//Standard Output
// 	res.Status = false
// 	//Check request data cannot be empty
// 	if request.Script == "" && request.SerialID == "" {
// 		res.Error = "Request Data Cannot Be Empty"
// 		c.JSON(http.StatusExpectationFailed, res)
// 		return
// 	}
// 	// res, err := executeScriptService(input)
// 	res, err = executeScriptService(request)
// 	if err != nil {
// 		logger.Log.Sugar().Errorf("Error executing script", err)
// 		// res.Error = "Error Executing Script Please Check"
// 		res.Error = err.Error()
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	// s := SanitizeScript(res)
// 	logger.Log.Info("OUT:executeScript")
// 	res.Status = true
// 	c.JSON(http.StatusOK, res)
// }
// func SanitizeScript(script string) string {
// 	s2 := strings.Replace(script, `\n`, "\n", -1)
// 	// s := strings.ReplaceAll(s2, "\\", "")
// 	return s2
// }
