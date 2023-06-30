package auth

import (
	"infraguard-manager/helpers/configHelper"
	"infraguard-manager/helpers/logger"
	model "infraguard-manager/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/robfig/cron"
)

var jwtKey string

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
	err := c.AddFunc("@every 50s", GenerateJWT)
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
