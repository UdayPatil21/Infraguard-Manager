package auth

import (
	"bytes"
	"encoding/json"
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/robfig/cron"
)

var (
	jwtKey     string
	LoginToken string
	UserToken  string
)

type JWTClaims struct {
	jwt.StandardClaims
}

func Init() {
	//Create token on server start
	GenerateJWT()
	//Schedule Refresh Token after every 24 hours
	//For new token in 24 hours
	RefreshToken()
}
func RefreshToken() {
	c := cron.New()
	err := c.AddFunc("0 0 * * *", GenerateJWT)
	if err != nil {
		logger.Info("Error", err)
	}
	c.Start()
}

//Generate JWT Tokens
func GenerateJWT() {
	jwtKey = configHelper.GetString("JwtSecretKey")

	//Create expiration time of JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenError := token.SignedString([]byte(jwtKey))
	model.TokenString = tokenString
	model.TokenError = tokenError

}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Isgmail  bool   `json:"isgmail"`
	Source   string `json:"source"`
}

type Data struct {
	Token        string `json:"token"`
	IsMFAEnabled string `json:"isMFAEnabled"`
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// type UserToken struct {
// 	Token        string `json:"token"`
// 	IsMFAEnabled string `json:"isMFAEnabled"`
// }

//Login to infraguard server and initiate authentication process
func InfraLogin() {

	loginDetails := Login{
		// email:    configHelper.GetString("email"),
		// password: configHelper.GetString("password"),
		Isgmail: false,
		Source:  "infraguard",
	}
	loginDetails.Email = configHelper.GetString("email")
	loginDetails.Password = configHelper.GetString("password")
	loginBytes, err := json.Marshal(loginDetails)
	if err != nil {
		logger.Error("Error marshling data login data", err)
		// return err
	}
	jsonStr := string(loginBytes)
	//Get server URL from config
	base_url := configHelper.GetString("Infraguard-URL")
	//create req add neccessary headers
	client := &http.Client{}
	req, _ := http.NewRequest("POST", base_url+"/api/auth/login", bytes.NewBufferString(jsonStr))
	// req.Header.Set("Authorization", configHelper.GetString("Authorization"))
	req.Header.Set("Access-Infraguard", configHelper.GetString("Access-Infraguard"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending agent data to infraguard server", err)
		// return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		// return err
	}
	logger.Info(string(respBytes))
	res := Response{}
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		logger.Error("Error unmarshlling login response", err)
	}
	LoginToken = res.Data.Token
	GenerateUserToken(res.Data.Token)
	defer resp.Body.Close()

}

//Get user token from infragurd server
//By using login token
func GenerateUserToken(token string) {
	body := Data{
		Token:        token,
		IsMFAEnabled: "No",
	}
	reqBytes, err := json.Marshal(body)
	if err != nil {
		logger.Error("Error marshling data login data", err)
		// return err
	}
	jsonStr := string(reqBytes)
	//Get server URL from config
	base_url := configHelper.GetString("Infraguard-URL")
	//create req add neccessary headers
	client := &http.Client{}
	req, _ := http.NewRequest("POST", base_url+"/api/auth/authverification", bytes.NewBufferString(jsonStr))
	// req.Header.Set("Authorization", configHelper.GetString("Authorization"))
	req.Header.Set("Access-Infraguard", configHelper.GetString("Access-Infraguard"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending agent data to infraguard server", err)
		// return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response", err)
		// return err
	}
	logger.Info(string(respBytes))
	res := Response{}
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		logger.Error("Error unmarshlling login response", err)
	}
	UserToken = res.Data.Token
}
