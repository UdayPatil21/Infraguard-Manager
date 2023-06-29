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

//Generate JWT Tokens
func GenerateJWT() {
	jwtKey = configHelper.GetString("JwtSecretKey")
	c := cron.New()
	err := c.AddFunc("@every 50s", func() {
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
	})
	if err != nil {
		logger.Info("Error", err)
	}
	c.Start()

}
