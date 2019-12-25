package model

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaimModel ..
type JWTClaimModel struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Group    string `json:"Group"`
}

// UserInfo ...
type UserInfo struct {
	Username string `json:"Username"`
	Group    string `json:"Group"`
}
