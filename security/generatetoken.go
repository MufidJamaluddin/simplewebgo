package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	model "github.com/MufidJamaluddin/simplewebgo/model"
	constant "github.com/MufidJamaluddin/simplewebgo/constant"
)

// GenerateJWTClaim ..
func generateJWTClaim(username string, group string) *model.JWTClaimModel {

	claim := &model.JWTClaimModel {
		StandardClaims: jwt.StandardClaims {
			Issuer:   constant.ApplicationName,
			ExpiresAt: time.Now().Add(constant.LoginExpirationDuration).Unix(),
		},
		Username: username,
		Group:    group,
	}

	return claim
}

// GenerateJWTToken ...
func GenerateJWTToken(username string, group string) *jwt.Token {
	refclaims := generateJWTClaim(username, group)
	token := jwt.NewWithClaims(
		constant.JWTSigningMethod,
		(*refclaims),
	)
	return token
}